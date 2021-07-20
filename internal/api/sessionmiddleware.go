package api

import (
	"encoding/hex"
	"github.com/SpaceXLaunchBot/site/internal/encryption"
	"github.com/gin-gonic/gin"
	"time"
)

// TODO: In middlewares, what's the best way of stopping? https://github.com/gin-gonic/gin/issues/853

// SessionMiddleware checks for and passes along a decrypted users session as a database.SessionRecord.
func (a Api) SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("sessionId")
		if err != nil {
			endWithResponse(c, responseNoSession)
			return
		}

		sessionKeyHex, err := c.Cookie("sessionKey")
		if err != nil {
			endWithResponse(c, responseNoSession)
			return
		}
		sessionKey, err := hex.DecodeString(sessionKeyHex)
		if err != nil {
			endWithResponse(c, responseInternalError)
			return
		}

		exists, session, err := a.db.GetSession(sessionId)
		if err != nil {
			endWithResponse(c, responseDatabaseError)
			return
		}
		if exists == false {
			endWithResponse(c, responseNoSession)
			return
		}

		// NOTE: We need to decrypt the stored information, the struct has room for the decrypted as well (below).
		accessTokenBytes, err := encryption.Decrypt(sessionKey, session.AccessTokenEncrypted)
		if err != nil {
			// NOTE: It doesn't matter if we fail to remove, we will be here when they make another request.
			_, _ = a.db.RemoveSession(sessionId)
			endWithResponse(c, responseNoSession)
			return
		}

		refreshTokenBytes, err := encryption.Decrypt(sessionKey, session.RefreshTokenEncrypted)
		if err != nil {
			_, _ = a.db.RemoveSession(sessionId)
			endWithResponse(c, responseNoSession)
			return
		}

		session.AccessToken = string(accessTokenBytes)
		session.RefreshToken = string(refreshTokenBytes)

		if session.AccessTokenExpiresAt.After(time.Now()) == false {
			// Everything is valid but our access token is expired.
			// TODO: Attempt to refresh with refresh token.
			endWithResponse(c, responseNoSession)
			return
		}

		c.Set("session", session)
		c.Next()
	}
}

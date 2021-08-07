// Helper functions for dealing with client sessions.

package api

import (
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/SpaceXLaunchBot/site/internal/encryption"
	"github.com/gin-gonic/gin"
)

// encryptAndSetTokens takes a discord.TokenResponse and inserts it, encrypted, into the
// database.
func (a Api) encryptAndSetTokens(sessionId string, sessionKey []byte, tokens discord.TokenResponse) error {
	accessTokenEncrypted, err := encryption.Encrypt(sessionKey, []byte(tokens.AccessToken))
	if err != nil {
		return err
	}

	refreshTokenEncrypted, err := encryption.Encrypt(sessionKey, []byte(tokens.RefreshToken))
	if err != nil {
		return err
	}

	changed, err := a.db.SetSession(sessionId, accessTokenEncrypted, tokens.ExpiresIn, refreshTokenEncrypted)
	if err != nil || !changed {
		return err
	}

	return nil
}

// removeSessionCookies removes the session cookies for the given context.
func (a Api) removeSessionCookies(c *gin.Context) {
	c.SetCookie("sessionId", "", 0, "/", a.hostName, a.isHTTPS, true)
	c.SetCookie("sessionKey", "", 0, "/", a.hostName, a.isHTTPS, true)
}

// endWithInvalidateSession is similar to endWithResponse but is used to invalidate the users session.
func (a Api) endWithInvalidateSession(c *gin.Context, id string) {
	// TODO: What happens if we remove cookies but fail to remove from db?
	_, _ = a.db.RemoveSession(id)
	a.removeSessionCookies(c)
	endWithResponse(c, responseNoSession)
}

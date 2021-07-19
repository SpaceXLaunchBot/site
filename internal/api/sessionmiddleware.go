package api

import (
	"context"
	"encoding/hex"
	"github.com/SpaceXLaunchBot/site/internal/encryption"
	"net/http"
	"time"
)

// SessionMiddleware defines a middleware to work with Gorilla/mux that checks for and passes along a users session.
func (a Api) SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIdCookie, err := r.Cookie("sessionId")
		if err != nil {
			endWithResponse(w, responseNoSession)
			return
		}
		sessionId := sessionIdCookie.Value

		sessionKeyCookie, err := r.Cookie("sessionKey")
		if err != nil {
			endWithResponse(w, responseNoSession)
			return
		}
		sessionKey, err := hex.DecodeString(sessionKeyCookie.Value)
		if err != nil {
			endWithResponse(w, responseInternalError)
			return
		}

		exists, session, err := a.db.GetSession(sessionId)
		if err != nil {
			endWithResponse(w, responseDatabaseError)
			return
		}
		if exists == false {
			endWithResponse(w, responseNoSession)
			return
		}

		// NOTE: We need to decrypt the stored information, the struct has room for the decrypted as well (below).
		accessTokenBytes, err := encryption.Decrypt(sessionKey, session.AccessTokenEncrypted)
		if err != nil {
			// NOTE: It doesn't matter if we fail to remove, we will be here when they make another request.
			_, _ = a.db.RemoveSession(sessionId)
			endWithResponse(w, responseNoSession)
			return
		}

		refreshTokenBytes, err := encryption.Decrypt(sessionKey, session.RefreshTokenEncrypted)
		if err != nil {
			_, _ = a.db.RemoveSession(sessionId)
			endWithResponse(w, responseNoSession)
			return
		}

		session.AccessToken = string(accessTokenBytes)
		session.RefreshToken = string(refreshTokenBytes)

		if session.AccessTokenExpiresAt.After(time.Now()) == false {
			// Everything is valid but our access token is expired.
			// TODO: Attempt to refresh with refresh token. Not sure where in codebase to do this.
			endWithResponse(w, responseNoSession)
			return
		}

		// https://github.com/julienschmidt/httprouter/issues/198#issuecomment-305784346
		ctx := r.Context()
		ctx = context.WithValue(ctx, "session", session)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

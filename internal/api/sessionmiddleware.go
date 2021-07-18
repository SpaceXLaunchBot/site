package api

import (
	"context"
	"encoding/hex"
	"net/http"
)

// SessionMiddleware defines a middleware to work with Gorilla/mux that checks for and passes along a users session.
func (a Api) SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("session")
		if err != nil {
			endWithResponse(w, responseNoSession)
			return
		}

		sessionKeyCookie, err := r.Cookie("key")
		if err != nil {
			endWithResponse(w, responseNoSession)
			return
		}
		sessionKey, err := hex.DecodeString(sessionKeyCookie.Value)
		if err != nil {
			endWithResponse(w, responseInternalError)
			return
		}

		exists, session, err := a.db.GetSession(sessionCookie.Value, sessionKey)
		if err != nil {
			endWithResponse(w, responseDatabaseError)
			return
		}
		if exists == false {
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

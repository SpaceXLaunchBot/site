package api

import (
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

// See this discussion as to why I chose to do it this way:
// https://security.stackexchange.com/a/77316

// New auth plan:
//  - User logs in to discord, returns to "/api/login" with a "code" query param
//  - Extract this param server-side, use it to request token(s) from Discord
//  - Receive access and refresh tokens
//  - Create a session id
//    - Put this id in cookie and add to response
//    - Save in db alongside access tokens and expiry time of access token
//  - Redirect to homepage, now they have a session
//  - Now when a user makes a request we look up their session and get their info

// TODO:
//  - Create an encryption key before saving stuff to db
//  - Give this key to client in a cookie
//  - Encrypt the tokens with this and store them in the db alongside another random value
//  - Now when a user makes a request we have to decrypt tokens with their cookie before use
//  - If their token fails we want to refresh it using their refresh token with the Discord OAuth API

// HandleDiscordLogin is the endpoint handler for when a user authenticates using Discords OAuth flow.
func (a Api) HandleDiscordLogin(w http.ResponseWriter, r *http.Request) {
	oAuthCode := r.URL.Query().Get("code")
	if oAuthCode == "" {
		// TODO: Handle this error and all errors in all functions in this file.
		return
	}

	tokens, _ := a.discordClient.TokensFromCode(oAuthCode)
	sessionId := uuid.NewV4().String()

	// TODO: As well as error what do we do if SetSession returns changed as false?
	_, _ = a.db.SetSession(sessionId, tokens.AccessToken, tokens.ExpiresIn, "")

	// Currently we don't store or use the refresh token, so the cookie expiry time can be the same as the access tokens.
	sessionCookie := &http.Cookie{
		Name:     "session",
		Value:    sessionId,
		Path:     "/",
		Domain:   a.hostName,
		MaxAge:   tokens.ExpiresIn,
		Secure:   a.isHTTPS,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, sessionCookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// HandleDiscordLogout is for when a user wants to invalidate their session that contains Discord OAuth info.
func (a Api) HandleDiscordLogout(w http.ResponseWriter, r *http.Request) {
	sessionCookie, _ := r.Cookie("session")
	// Erase the servers knowledge of it.
	_, _ = a.db.RemoveSession(sessionCookie.Value)
	// Erase the client side session cookie by removing its value and set expiry time to 0.
	c := &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		Domain:   a.hostName,
		Secure:   a.isHTTPS,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

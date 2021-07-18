package api

import (
	"encoding/json"
	"github.com/SpaceXLaunchBot/site/internal/database"
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

type discordLoginJson struct {
	Code string `json:"code"`
}

// HandleDiscordLogin is the endpoint handler for when a user authenticates using Discords OAuth flow.
func (a Api) HandleDiscordLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var loginInfo discordLoginJson
	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		endWithResponse(w, responseBadJson)
		return
	}

	tokens, err := a.discordClient.TokensFromCode(loginInfo.Code)
	if err != nil {
		endWithResponse(w, responseInvalidOAuthCode)
		return
	}

	sessionId := uuid.NewV4().String()

	changed, err := a.db.SetSession(sessionId, tokens.AccessToken, tokens.ExpiresIn, "")
	if err != nil || !changed {
		endWithResponse(w, responseDatabaseError)
		return
	}

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
	endWithResponse(w, responseAllOk)
}

// HandleDiscordLogout is for when a user wants to invalidate their session that contains Discord OAuth info.
func (a Api) HandleDiscordLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	session := r.Context().Value("session").(database.SessionRecord)

	// Erase the servers knowledge of it.
	deleted, err := a.db.RemoveSession(session.Session)
	if err != nil || !deleted {
		endWithResponse(w, responseDatabaseError)
		return
	}

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
	endWithResponse(w, responseAllOk)
}

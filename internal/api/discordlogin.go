package api

import (
	"encoding/hex"
	"encoding/json"
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/encryption"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

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
	if loginInfo.Code == "" {
		endWithResponse(w, responseInvalidOAuthCode)
		return
	}

	tokens, err := a.discordClient.TokensFromCode(loginInfo.Code)
	if err != nil {
		endWithResponse(w, responseInvalidOAuthCode)
		return
	}

	sessionId := uuid.NewV4().String()
	sessionKey, err := encryption.GenerateKey()
	if err != nil {
		endWithResponse(w, responseInternalError)
		return
	}

	changed, err := a.db.SetSession(sessionId, sessionKey, tokens.AccessToken, tokens.ExpiresIn, tokens.RefreshToken)
	if err != nil || !changed {
		endWithResponse(w, responseDatabaseError)
		return
	}

	// Currently we don't use the refresh token, so the cookie expiry time can be the same as the access tokens.
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

	sessionKeyHex := hex.EncodeToString(sessionKey)
	keyCookie := &http.Cookie{
		Name:     "key",
		Value:    sessionKeyHex,
		Path:     "/",
		Domain:   a.hostName,
		MaxAge:   tokens.ExpiresIn,
		Secure:   a.isHTTPS,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, sessionCookie)
	http.SetCookie(w, keyCookie)
	endWithResponse(w, responseAllOk)
}

// HandleDiscordLogout is for when a user wants to invalidate their session that contains Discord OAuth info.
func (a Api) HandleDiscordLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	session := r.Context().Value("session").(database.SessionRecord)

	// Erase the servers knowledge of it.
	deleted, err := a.db.RemoveSession(session.SessionID)
	if err != nil || !deleted {
		endWithResponse(w, responseDatabaseError)
		return
	}

	// Erase the client side cookies by removing their values and setting expiry times to 0.
	sessionCookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		Domain:   a.hostName,
		Secure:   a.isHTTPS,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0),
	}

	keyCookie := &http.Cookie{
		Name:     "key",
		Value:    "",
		Path:     "/",
		Domain:   a.hostName,
		Secure:   a.isHTTPS,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, sessionCookie)
	http.SetCookie(w, keyCookie)
	endWithResponse(w, responseAllOk)
}

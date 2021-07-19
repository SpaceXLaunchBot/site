package api

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/encryption"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

// These are only to be used by the login handler, as they will be returned in query params to "/".
var errInvalidOauthCode = errors.New("invalid OAuth code")
var errCryptoKeyGenFailed = errors.New("server failed to generate encryption key for secrets")
var errEncryptionFailed = errors.New("the server failed to encrypt your secrets")
var errDatabaseErr = errors.New("database error")

type discordLoginJson struct {
	Code string `json:"code"`
}

func loginError(w http.ResponseWriter, r *http.Request, err error) {
	// Try to pass in a custom error rather than an unknown one to help prevent possible sensitive data leaks.
	http.Redirect(w, r, fmt.Sprintf("/?login_error=%s", err.Error()), http.StatusSeeOther)
}

// HandleDiscordLogin is the endpoint handler for when a user authenticates using Discords OAuth flow.
func (a Api) HandleDiscordLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	oauthCode := r.URL.Query().Get("code")
	if oauthCode == "" {
		loginError(w, r, errInvalidOauthCode)
		return
	}

	tokens, err := a.discordClient.TokensFromCode(oauthCode)
	if err != nil {
		loginError(w, r, errInvalidOauthCode)
		return
	}

	sessionId := uuid.NewV4().String()
	sessionKey, err := encryption.GenerateKey()
	if err != nil {
		loginError(w, r, errCryptoKeyGenFailed)
		return
	}

	accessTokenEncrypted, err := encryption.Encrypt(sessionKey, []byte(tokens.AccessToken))
	if err != nil {
		loginError(w, r, errEncryptionFailed)
		return
	}

	refreshTokenEncrypted, err := encryption.Encrypt(sessionKey, []byte(tokens.RefreshToken))
	if err != nil {
		loginError(w, r, errEncryptionFailed)
		return
	}

	changed, err := a.db.SetSession(sessionId, accessTokenEncrypted, tokens.ExpiresIn, refreshTokenEncrypted)
	if err != nil || !changed {
		loginError(w, r, errDatabaseErr)
		return
	}

	// Currently we don't use the refresh token, so the cookie expiry time can be the same as the access tokens.
	sessionCookie := &http.Cookie{
		Name:     "sessionId",
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
		Name:     "sessionKey",
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// HandleDiscordLogout is for when a user wants to invalidate their session that contains Discord OAuth info.
func (a Api) HandleDiscordLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	session := r.Context().Value("session").(database.SessionRecord)

	// Erase the servers knowledge of it.
	deleted, err := a.db.RemoveSession(session.SessionId)
	if err != nil || !deleted {
		endWithResponse(w, responseDatabaseError)
		return
	}

	// Erase the client side cookies by removing their values and setting expiry times to 0.
	sessionCookie := &http.Cookie{
		Name:     "sessionId",
		Value:    "",
		Path:     "/",
		Domain:   a.hostName,
		Secure:   a.isHTTPS,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0),
	}

	keyCookie := &http.Cookie{
		Name:     "sessionKey",
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

package api

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"

	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/encryption"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// These errors are only to be used by HandleDiscordLogin, as they will be returned to
// the user as plaintext in the redirect query params.
var errNoOauthCode = errors.New("no OAuth code query parameter")
var errInvalidOauthCode = errors.New("invalid OAuth code")
var errCryptoKeyGenFailed = errors.New("server failed to generate encryption key for secrets")

// var errEncryptionFailed = errors.New("the server failed to encrypt your secrets")
// var errDatabaseErr = errors.New("database error")

// endWithLoginError is similar to endWithResponse but redirects instead of sending JSON.
func endWithLoginError(c *gin.Context, err error) {
	// Try to pass in a custom error rather than an unknown one to help prevent possible sensitive data leaks.
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?login_error=%s", err.Error()))
}

// HandleDiscordLogin is the endpoint handler for when a user authenticates using Discords OAuth flow.
func (a Api) HandleDiscordLogin(c *gin.Context) {
	oauthCode := c.Request.URL.Query().Get("code")
	if oauthCode == "" {
		endWithLoginError(c, errNoOauthCode)
		return
	}

	tokens, err := a.discordClient.TokensFromCode(oauthCode)
	if err != nil {
		endWithLoginError(c, errInvalidOauthCode)
		return
	}

	sessionId := uuid.NewV4().String()
	sessionKey, err := encryption.GenerateKey()
	if err != nil {
		endWithLoginError(c, errCryptoKeyGenFailed)
		return
	}

	// TODO: ASAP: Make sure err is user friendly and non leaky.
	if err = a.encryptAndSetTokens(sessionId, sessionKey, tokens); err != nil {
		endWithLoginError(c, err)
		return
	}

	// Currently we don't use the refresh token, so the cookie expiry time can be the same as the access tokens.
	c.SetCookie("sessionId", sessionId, tokens.ExpiresIn, "/", a.hostName, a.isHTTPS, true)

	sessionKeyHex := hex.EncodeToString(sessionKey)
	c.SetCookie("sessionKey", sessionKeyHex, tokens.ExpiresIn, "/", a.hostName, a.isHTTPS, true)

	c.Redirect(http.StatusSeeOther, "/")
}

// HandleDiscordLogout is for when a user wants to invalidate their session that contains Discord OAuth info.
func (a Api) HandleDiscordLogout(c *gin.Context) {
	session := c.MustGet("session").(database.SessionRecord)

	// Erase the servers knowledge of it.
	deleted, err := a.db.RemoveSession(session.SessionId)
	if err != nil || !deleted {
		endWithResponse(c, responseDatabaseError)
		return
	}

	a.removeSessionCookies(c)
	endWithResponse(c, responseAllOk)
}
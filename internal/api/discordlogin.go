package api

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/encryption"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
)

// These are only to be used by the login handler, as they will be returned in the redirect query params.
var errNoOauthCode = errors.New("no OAuth code query parameter")
var errInvalidOauthCode = errors.New("invalid OAuth code")
var errCryptoKeyGenFailed = errors.New("server failed to generate encryption key for secrets")
var errEncryptionFailed = errors.New("the server failed to encrypt your secrets")
var errDatabaseErr = errors.New("database error")

func loginError(c *gin.Context, err error) {
	// Try to pass in a custom error rather than an unknown one to help prevent possible sensitive data leaks.
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?login_error=%s", err.Error()))
}

// HandleDiscordLogin is the endpoint handler for when a user authenticates using Discords OAuth flow.
func (a Api) HandleDiscordLogin(c *gin.Context) {
	oauthCode := c.Request.URL.Query().Get("code")
	if oauthCode == "" {
		loginError(c, errNoOauthCode)
		return
	}

	tokens, err := a.discordClient.TokensFromCode(oauthCode)
	if err != nil {
		loginError(c, errInvalidOauthCode)
		return
	}

	sessionId := uuid.NewV4().String()
	sessionKey, err := encryption.GenerateKey()
	if err != nil {
		loginError(c, errCryptoKeyGenFailed)
		return
	}

	accessTokenEncrypted, err := encryption.Encrypt(sessionKey, []byte(tokens.AccessToken))
	if err != nil {
		loginError(c, errEncryptionFailed)
		return
	}

	refreshTokenEncrypted, err := encryption.Encrypt(sessionKey, []byte(tokens.RefreshToken))
	if err != nil {
		loginError(c, errEncryptionFailed)
		return
	}

	changed, err := a.db.SetSession(sessionId, accessTokenEncrypted, tokens.ExpiresIn, refreshTokenEncrypted)
	if err != nil || !changed {
		loginError(c, errDatabaseErr)
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

	c.SetCookie("sessionId", "", 0, "/", a.hostName, a.isHTTPS, true)
	c.SetCookie("sessionKey", "", 0, "/", a.hostName, a.isHTTPS, true)
	endWithResponse(c, responseAllOk)
}

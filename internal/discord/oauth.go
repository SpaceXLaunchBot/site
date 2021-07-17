package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var tokenUrl = "https://discord.com/api/oauth2/token"

// TokenResponse is to marshal the Discord OAuth token response into.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// TokensFromCode takes an OAuth code provided by Discord and fetches access and refresh tokens.
func (c Client) TokensFromCode(code string) (TokenResponse, error) {
	payload := strings.NewReader(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=authorization_code&code=%s&redirect_uri=%s",
		c.clientId, c.clientSecret, code, c.redirectUri,
	))

	req, err := http.NewRequest("POST", tokenUrl, payload)
	if err != nil {
		return TokenResponse{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return TokenResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return TokenResponse{}, err
	}

	tr := TokenResponse{}
	err = json.Unmarshal(body, &tr)
	return tr, err
}

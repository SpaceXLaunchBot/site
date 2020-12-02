package discord

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// TODO: Deal with being rate limited by the Discord API. Currently unclear as to if the rate limit headers are s or ms.

const apiBase = "https://discord.com/api"

// ErrBadAuth described an error that occurs when a request fails due to bad authorization.
var ErrBadAuth = errors.New("failed to authorize with Discord API")

// ErrRateLimit describes an error that occurs when the discord.Client is getting rate limited.
var ErrRateLimit = errors.New("hit Discord API rate limit, try again in a few seconds")

// apiRequest performs a DiscordClient API request.
func apiRequest(endpoint, bearerToken string) ([]byte, error) {
	apiUrl, _ := url.Parse(apiBase)
	apiUrl.Path = path.Join(apiUrl.Path, endpoint)

	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", bearerToken))

	c := http.Client{Timeout: time.Second * 10}
	res, err := c.Do(req)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf("request to Discord API failed: %s", err))
	}

	if res.StatusCode == 401 {
		return []byte{}, ErrBadAuth
	}
	if res.StatusCode == 429 {
		return []byte{}, ErrRateLimit
	}

	body, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

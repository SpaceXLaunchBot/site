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

const apiBase = "https://discord.com/api"

// ErrBadAuth described an error that occurs when a request fails due to bad authorization.
var ErrBadAuth = errors.New("oauth authorization failed")

// apiRequest performs a Discord API request.
func apiRequest(endpoint, bearerToken string) ([]byte, error) {
	client := http.Client{Timeout: time.Second * 10}
	apiUrl, _ := url.Parse(apiBase)
	apiUrl.Path = path.Join(apiUrl.Path, endpoint)

	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", bearerToken))
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode == 401 {
		return []byte{}, ErrBadAuth
	}

	body, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

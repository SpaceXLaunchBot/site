package discord

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

const apiBase = "https://discord.com/api"

// ErrBadAuth described an error that occurs when a request fails due to bad authorization.
var ErrBadAuth = errors.New("failed to authorize with the Discord API")

// ErrRateLimit describes an error that occurs when the discord.Client is getting rate limited.
var ErrRateLimit = errors.New("hit the Discord API rate limit, try again in a few seconds")

// Client contains methods for interacting with the Discord API.
type Client struct {
	httpClient *http.Client
	// NOTE: Each Discord API endpoint & each bearer token has their own rate limits (along with a global rate limit).
	//  A cache is used so that if a certain user is spamming a particular route, we won't spam the Discord API and
	//  constantly hit a rate limit. This won't save us from hitting the rate limit if we are getting lots of new
	//  different users (i.e. multiple per second), but I think that is unlikely.
	//  This does mean that the data could be out of date by a few seconds, e.g. if a user is no longer an admin they
	//  might still be able to make changes, but I think that's unlikely and also a fine tradeoff for not getting
	//  banned from the Discord API.
	cache *cache.Cache
}

// NewClient creates a new Client.
func NewClient() Client {
	return Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		cache:      cache.New(10*time.Second, 10*time.Minute),
	}
}

// apiRequest performs a DiscordClient API request.
func (c Client) apiRequest(endpoint, bearerToken string) ([]byte, error) {
	cacheKey := endpoint + bearerToken

	if cached, ok := c.cache.Get(cacheKey); ok {
		return cached.([]byte), nil
	}

	apiUrl, _ := url.Parse(apiBase)
	apiUrl.Path = path.Join(apiUrl.Path, endpoint)

	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
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

	c.cache.Set(cacheKey, body, cache.DefaultExpiration)
	return body, nil
}

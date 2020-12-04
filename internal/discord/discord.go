package discord

import (
	"errors"
	"fmt"
	"github.com/psidex/SpaceXLaunchBotSite/internal/cache"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

const apiBase = "https://discord.com/api"

// ErrBadAuth described an error that occurs when a request fails due to bad authorization.
var ErrBadAuth = errors.New("failed to authorize with Discord API")

// ErrRateLimit describes an error that occurs when the discord.Client is getting rate limited.
var ErrRateLimit = errors.New("hit Discord API rate limit, try again in a few seconds")

// Client contains methods for interacting with the Discord API.
type Client struct {
	httpClient *http.Client
	// NOTE: Each Discord API endpoint has it's own rate limits so have caches for each endpoint.
	// A cache like this is used so that if a user is spamming refresh or save or whatever, we won't spam the Discord
	// API and constantly hit a rate limit. This won't save us from hitting the rate limit if we are getting lots of
	// new different users (i.e. multiple per second), but I think that is unlikely.
	// TODO: I don't actually know if the above it 100% true, rate limits might also be different per auth token?
	// This does mean that the data could be out of date by the lifespan of the cache, e.g. if a user is no longer an
	// admin they might still be able to make changes, but I think that's unlikely and also a fine tradeoff for not
	// getting banned from the Discord API.
	guildListCache *cache.TimedCache
}

// NewClient creates a new Client.
func NewClient(httpClientTimeout time.Duration, cacheLifespan time.Duration) Client {
	return Client{
		httpClient:     &http.Client{Timeout: httpClientTimeout},
		guildListCache: cache.NewTimedCache(cacheLifespan.Seconds()),
	}
}

// apiRequest performs a DiscordClient API request.
func (c Client) apiRequest(endpoint, bearerToken string) ([]byte, error) {
	apiUrl, _ := url.Parse(apiBase)
	apiUrl.Path = path.Join(apiUrl.Path, endpoint)

	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", bearerToken))

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
	return body, nil
}
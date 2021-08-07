package api

import (
	"time"

	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/patrickmn/go-cache"
)

// Api contains everything required to run the API section of the server.
type Api struct {
	db            database.Db
	cache         *cache.Cache
	discordClient discord.Client
	// Host URI info for cookies.
	hostName string
	isHTTPS  bool
}

// NewApi creates a new Api.
func NewApi(db database.Db, client discord.Client, hostName, protocol string) Api {
	secure := true
	if protocol == "http:" {
		secure = false
	}
	return Api{
		db:            db,
		discordClient: client,
		cache:         cache.New(10*time.Second, 10*time.Minute),
		hostName:      hostName,
		isHTTPS:       secure,
	}
}

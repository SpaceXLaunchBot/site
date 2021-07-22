package api

import (
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"time"
)

// Api contains methods that interface with the database and the Discord API through a REST API.
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

// GuildListMiddleware gets a GuildList using the clients session.
func (a Api) GuildListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.MustGet("session").(database.SessionRecord)

		guilds, err := a.discordClient.GetGuildList(session.AccessToken)
		if err != nil {
			if err == discord.ErrBadAuth {
				a.endWithInvalidateSession(c, session.SessionId)
				return
			}
			resp := responseDiscordApiError
			// Add context to general error message.
			resp.Error += err.Error()
			endWithResponse(c, resp)
			return
		}

		c.Set("guilds", guilds)
		c.Next()
	}
}

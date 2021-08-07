package api

import (
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/gin-gonic/gin"
)

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

package api

import (
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/gin-gonic/gin"
)

// deleteChannelJson is a struct to marshal the api request data into.
type deleteChannelJson struct {
	ID      string `json:"id"`
	GuildID string `json:"guild_id"`
}

// DeleteChannel deletes ("unsubscribes") a channel from the database.
func (a Api) DeleteChannel(c *gin.Context) {
	guilds := c.MustGet("guilds").(discord.GuildList)

	var requestedDelete deleteChannelJson
	if err := c.ShouldBind(&requestedDelete); err != nil {
		endWithResponse(c, responseBadJson)
		return
	}

	allowedToEdit := false
	for _, guild := range guilds {
		if guild.HasAdminPerms() && guild.ID == requestedDelete.GuildID {
			allowedToEdit = true
		}
	}
	if !allowedToEdit {
		endWithResponse(c, responseNotAdmin)
		return
	}

	changed, err := a.db.DeleteSubscribedChannel(
		requestedDelete.ID,
		requestedDelete.GuildID,
	)
	if err != nil {
		endWithResponse(c, responseDatabaseError)
		return
	}
	if !changed {
		endWithResponse(c, responseChannelNotInGuild)
		return
	}

	endWithResponse(c, responseAllOk)
}

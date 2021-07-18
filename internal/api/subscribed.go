package api

import (
	"fmt"
	"github.com/SpaceXLaunchBot/site/internal/database"
	"net/http"
)

// subscribedResponse is the API response for the subscribed channels API route.
type subscribedResponse struct {
	genericResponse
	Subscribed map[string]*guildDetails `json:"subscribed"`
}

// guildDetails holds information about a guild.
type guildDetails struct {
	Name               string                             `json:"name"`
	Icon               string                             `json:"icon"`
	SubscribedChannels []database.SubscribedChannelRecord `json:"subscribed_channels"`
}

// SubscribedChannels returns a list of information about guilds user is authed in that are subscribed to the notification service.
func (a Api) SubscribedChannels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	guilds, sentErr := a.getGuildList(w, r)
	if sentErr == true {
		return
	}

	// TODO: non-branch: 3 data structures is ez but possibly not the most efficient.
	var adminGuilds []string
	adminGuildNames := make(map[string]string)
	adminGuildIcons := make(map[string]string)

	for _, guild := range guilds {
		if guild.HasAdminPerms() {
			adminGuilds = append(adminGuilds, guild.ID)
			adminGuildNames[guild.ID] = guild.Name
			adminGuildIcons[guild.ID] = fmt.Sprintf("https://cdn.discordapp.com/icons/%s/%s.png", guild.ID, guild.Icon)
		}
	}

	if len(adminGuilds) == 0 {
		endWithResponse(w, responseNotAdminInAny)
		return
	}

	subbedChannels, err := a.db.SubscribedChannels(adminGuilds)
	if err != nil {
		endWithResponse(w, responseDatabaseError)
		return
	}
	if len(subbedChannels) == 0 {
		endWithResponse(w, responseNoSubscribedInAny)
		return
	}

	resp := subscribedResponse{}
	resp.Success = true
	resp.Subscribed = make(map[string]*guildDetails)

	nonNilStr := ""

	for _, channel := range subbedChannels {
		if channel.LaunchMentions == nil {
			// If the pointer is nil point it to an empty string.
			channel.LaunchMentions = &nonNilStr
		}

		if details, ok := resp.Subscribed[channel.GuildId]; ok {
			details.SubscribedChannels = append(details.SubscribedChannels, channel)
		} else {
			resp.Subscribed[channel.GuildId] = &guildDetails{
				Name:               adminGuildNames[channel.GuildId],
				Icon:               adminGuildIcons[channel.GuildId],
				SubscribedChannels: []database.SubscribedChannelRecord{channel},
			}
		}
	}

	endWithResponse(w, &resp)
}

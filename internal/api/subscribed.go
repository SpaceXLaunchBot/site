package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// subscribedChannel holds information about a subscribed channel.
type subscribedChannel struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	NotificationType string `json:"notification_type"`
	LaunchMentions   string `json:"launch_mentions"`
}

// guildDetails holds information about a guild.
type guildDetails struct {
	Name               string              `json:"name"`
	Icon               string              `json:"icon"`
	SubscribedChannels []subscribedChannel `json:"subscribed_channels"`
}

// SubscribedChannelsResponse represents the JSON response from the SubscribedChannels endpoint.
type SubscribedChannelsResponse map[string]*guildDetails

// SubscribedChannels returns a list of information about guilds user is authed in that are subscribed to the notification service.
func (a Api) SubscribedChannels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	guilds, sentErr := a.getGuildList(w, r)
	if sentErr == true {
		return
	}

	// TODO: 3 data structures is ez but possibly not the most efficient.
	var adminGuilds []string
	adminGuildNames := make(map[string]string)
	adminGuildIcons := make(map[string]string)

	for _, guild := range guilds {
		if guild.HasAdminPerms() {
			adminGuilds = append(adminGuilds, guild.ID)
			adminGuildNames[guild.ID] = guild.Name

			iconUrl := fmt.Sprintf("https://cdn.discordapp.com/icons/%s/%s.png", guild.ID, guild.Icon)
			adminGuildIcons[guild.ID] = iconUrl
		}
	}

	if len(adminGuilds) == 0 {
		// I think this is the right status code for this sort of error.
		w.WriteHeader(http.StatusConflict)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "you do not have admin permissions in any guilds"})
		return
	}

	// What will become out API response, {guild id : guild details}.
	details := make(SubscribedChannelsResponse)

	subbedChannels, wClosed := a.db.SubscribedChannels(adminGuilds)
	if wClosed != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "database error :("})
		return
	}
	if len(subbedChannels) == 0 {
		w.WriteHeader(http.StatusConflict)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "you do not have any subscribed channels in guilds that you administrate"})
		return
	}

	for _, channel := range subbedChannels {
		channelStruct := subscribedChannel{
			Id:               channel.Id,
			Name:             channel.Name,
			NotificationType: channel.NotificationType,
			LaunchMentions:   channel.LaunchMentions.String,
		}

		if d, ok := details[channel.GuildId]; ok {
			d.SubscribedChannels = append(d.SubscribedChannels, channelStruct)
		} else {
			details[channel.GuildId] = &guildDetails{
				Name:               adminGuildNames[channel.GuildId],
				Icon:               adminGuildIcons[channel.GuildId],
				SubscribedChannels: []subscribedChannel{channelStruct},
			}
		}

	}

	w.WriteHeader(http.StatusOK)
	wClosed = json.NewEncoder(w).Encode(details)
	if wClosed != nil {
		log.Printf("Failed to encode SubscribedChannels response: %s", wClosed)
	}
}

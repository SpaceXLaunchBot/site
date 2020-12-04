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

// guildsWithSubscribedResponse represents the JSON response from the GuildsWithSubscribed endpoint.
type guildsWithSubscribedResponse map[string]*guildDetails

// GuildsWithSubscribed takes a discord oauth token and returns a list of information about guilds and channels that
// the user is in that are subscribed to the notification service.
func (a Api) GuildsWithSubscribed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	token := r.Header.Get("Discord-Bearer-Token")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "no Discord-Bearer-Token header"})
		return
	}

	guilds, err := a.discordClient.GetGuildList(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "error getting information from Discord API"})
		return
	}

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
	details := make(guildsWithSubscribedResponse)

	subbedChannels, err := a.db.SubscribedChannels(adminGuilds)
	if err != nil {
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
	err = json.NewEncoder(w).Encode(details)
	if err != nil {
		log.Printf("Failed to encode GuildsWithSubscribed response: %s", err)
	}
}

package api

import (
	"encoding/json"
	"fmt"
	"github.com/psidex/SpaceXLaunchBotSite/internal/database"
	"github.com/psidex/SpaceXLaunchBotSite/internal/discord"
	"log"
	"net/http"
)

// guildDetails is part of the API response for GuildsWithSubscribed and holds information about a guild.
type guildDetails struct {
	Name               string                       `json:"name"`
	Icon               string                       `json:"icon"`
	SubscribedChannels []database.SubscribedChannel `json:"subscribed_channels"`
}

// GuildsWithSubscribed takes a discord oauth token and returns a list of information about guilds and channels that
// the user is in that are subscribed to the notification service.
func (a Api) GuildsWithSubscribed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	token := r.Header.Get("Discord-Bearer-Token")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiError{Error: "no Discord-Bearer-Token header"})
		return
	}

	guilds, err := discord.GetGuildList(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiError{Error: err.Error()})
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
		_ = json.NewEncoder(w).Encode(apiError{Error: "you do not have admin permissions in any guilds"})
		return
	}

	// What will become out API response, {guild id : guild details}.
	details := make(map[string]*guildDetails)

	subbedChannels, err := a.db.SubscribedChannels(adminGuilds)
	if err != nil {
		// TODO: Returning the err from a db method might expose internal db stuff accidentally?.
		//  They also probably won't be user friendly.
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(apiError{Error: err.Error()})
		return
	}
	if len(subbedChannels) == 0 {
		w.WriteHeader(http.StatusConflict)
		_ = json.NewEncoder(w).Encode(apiError{Error: "you do not have any subscribed channels in guilds that you administrate"})
		return
	}

	for _, channel := range subbedChannels {
		if d, ok := details[channel.GuildId]; ok {
			d.SubscribedChannels = append(d.SubscribedChannels, channel)
		} else {
			details[channel.GuildId] = &guildDetails{
				Name:               adminGuildNames[channel.GuildId],
				Icon:               adminGuildIcons[channel.GuildId],
				SubscribedChannels: []database.SubscribedChannel{channel},
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(details)
	if err != nil {
		log.Printf("Failed to encode GuildsWithSubscribed response: %s", err)
	}
}

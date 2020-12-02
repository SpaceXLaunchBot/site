package api

import (
	"encoding/json"
	"github.com/psidex/SpaceXLaunchBotSite/internal/database"
	"github.com/psidex/SpaceXLaunchBotSite/internal/discord"
	"log"
	"net/http"
)

// Api contains methods that interface with the database through a REST API.
type Api struct {
	d database.Db
}

// NewApi creates a new Api.
func NewApi(d database.Db) Api {
	return Api{d: d}
}

// guildDetails is part of the API response for GuildsWithSubscribed and holds information about a guild.
type guildDetails struct {
	Name               string                       `json:"name"`
	SubscribedChannels []database.SubscribedChannel `json:"subscribed_channels"`
}

// GuildsWithSubscribed takes a discord oauth token and returns a list of information about guilds and channels that
// the user is in that are subscribed to the notification service.
func (a Api) GuildsWithSubscribed(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Discord-Bearer-Token")
	if token == "" {
		http.Error(w, "no Discord-Bearer-Token", http.StatusBadRequest)
		return
	}

	guilds, err := discord.GetGuildList(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var adminGuilds []string
	adminGuildNames := make(map[string]string)

	for _, guild := range guilds {
		// Admin is 0x00000008: https://discord.com/developers/docs/topics/permissions
		if 8&guild.Permissions != 0 {
			adminGuilds = append(adminGuilds, guild.ID)
			adminGuildNames[guild.ID] = guild.Name
		}
	}

	// What will become out API response, {guild id : guild details}.
	details := make(map[string]*guildDetails)

	// TODO: Actually inform client of err instead of doing nothing. Check other errs as well.
	subbedChannels, err := a.d.SubscribedChannels(adminGuilds)
	if err != nil {
		log.Printf("SubscribedChannels(adminGuilds) failed: %s", err)
		return
	}

	for _, channel := range subbedChannels {
		if d, ok := details[channel.GuildId]; ok {
			d.SubscribedChannels = append(d.SubscribedChannels, channel)
		} else {
			details[channel.GuildId] = &guildDetails{
				Name:               adminGuildNames[channel.GuildId],
				SubscribedChannels: []database.SubscribedChannel{channel},
			}
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(details)
	if err != nil {
		log.Printf("Encode(details) failed: %s", err)
	}
}

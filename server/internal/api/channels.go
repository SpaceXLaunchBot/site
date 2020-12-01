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

// guildDetails is the API response that contains information about guilds and subscribed channels for a user.
type guildDetails struct {
	Name               string             `json:"name"`
	ID                 string             `json:"id"`
	SubscribedChannels []database.Channel `json:"subscribed_channels"`
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

	// For each admin guild, get channels that are subbed for that guild.
	var details []guildDetails
	for _, guild := range guilds {
		// Admin is 0x00000008: https://discord.com/developers/docs/topics/permissions
		if 8&guild.Permissions != 0 {
			// TODO: What happens if err?
			subbed, _ := a.d.SubscribedChannels(guild.ID)
			if len(subbed) > 0 {
				details = append(details, guildDetails{
					Name:               guild.Name,
					ID:                 guild.ID,
					SubscribedChannels: subbed,
				})
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

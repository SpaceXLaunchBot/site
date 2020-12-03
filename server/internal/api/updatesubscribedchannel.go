package api

import (
	"encoding/json"
	"net/http"
)

// UpdateSubscribedChannel is NOT IMPLEMENTED YET.
func (a Api) UpdateSubscribedChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	token := r.Header.Get("Discord-Bearer-Token")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiError{Error: "no Discord-Bearer-Token header"})
		return
	}

	guilds, err := a.discordClient.GetGuildList(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiError{Error: err.Error()})
		return
	}

	var adminGuilds []string

	for _, guild := range guilds {
		if guild.HasAdminPerms() {
			adminGuilds = append(adminGuilds, guild.ID)
		}
	}
}

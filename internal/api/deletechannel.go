package api

import (
	"encoding/json"
	"net/http"
)

// deleteChannelJson is a struct to marshal the api request data into.
type deleteChannelJson struct {
	ID      string `json:"id"`
	GuildID string `json:"guild_id"`
}

// DeleteChannel deletes ("unsubscribes") a channel from the database.
func (a Api) DeleteChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	guilds, sentErr := a.getGuildList(w, r)
	if sentErr == true {
		return
	}

	var requestedDelete deleteChannelJson
	err := json.NewDecoder(r.Body).Decode(&requestedDelete)
	if err != nil {
		endWithResponse(w, responseBadJson)
		return
	}

	allowedToEdit := false
	for _, guild := range guilds {
		if guild.HasAdminPerms() && guild.ID == requestedDelete.GuildID {
			allowedToEdit = true
		}
	}
	if !allowedToEdit {
		endWithResponse(w, responseNotAdmin)
		return
	}

	changed, err := a.db.DeleteSubscribedChannel(
		requestedDelete.ID,
		requestedDelete.GuildID,
	)
	if err != nil {
		endWithResponse(w, responseDatabaseError)
		return
	}
	if !changed {
		endWithResponse(w, responseChannelNotInGuild)
		return
	}

	endWithResponse(w, responseAllOk)
}

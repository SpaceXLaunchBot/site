package api

import (
	"encoding/json"
	"net/http"
)

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
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "failed to decode JSON body"})
		return
	}

	allowedToEdit := false
	for _, guild := range guilds {
		if guild.HasAdminPerms() && guild.ID == requestedDelete.GuildID {
			allowedToEdit = true
		}
	}
	if !allowedToEdit {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "you are not an admin in that server"})
		return
	}

	changed, err := a.db.DeleteSubscribedChannel(
		requestedDelete.ID,
		requestedDelete.GuildID,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "database error :("})
		return
	}
	if !changed {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "no channel with that ID in the given guild"})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(apiResponse{Success: true})
}

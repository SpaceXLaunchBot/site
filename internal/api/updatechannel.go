package api

import (
	"encoding/json"
	"net/http"
)

type updateChannelJson struct {
	ID               string `json:"id"`
	GuildID          string `json:"guild_id"`
	NotificationType string `json:"notification_type"`
	LaunchMentions   string `json:"launch_mentions"`
}

// UpdateChannel updates information about a channel in the database.
func (a Api) UpdateChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	guilds, sentErr := a.getGuildList(w, r)
	if sentErr == true {
		return
	}

	var requestedUpdate updateChannelJson
	err := json.NewDecoder(r.Body).Decode(&requestedUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "failed to decode JSON body"})
		return
	}

	allowedToEdit := false
	for _, guild := range guilds {
		if guild.HasAdminPerms() && guild.ID == requestedUpdate.GuildID {
			allowedToEdit = true
		}
	}
	if !allowedToEdit {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "you are not an admin in that server"})
		return
	}

	// NOTE: The guild ID is required here to prevent someone passing the above checks (i.e. they are an admin in the
	//  provided guild) and then being able to edit a channel from another guild. We could check this by first querying
	//  the database for the guild ID given the channel ID, but this would be uselessly slower when the client already
	//  knows the guild ID.
	changed, err := a.db.UpdateSubscribedChannel(
		requestedUpdate.ID,
		requestedUpdate.GuildID,
		requestedUpdate.NotificationType,
		requestedUpdate.LaunchMentions,
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

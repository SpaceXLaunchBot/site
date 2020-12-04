package api

import (
	"encoding/json"
	"net/http"
)

type updateSubscribedChannelJson struct {
	ID               string `json:"id"`
	GuildID          string `json:"guild_id"`
	NotificationType string `json:"notification_type"`
	LaunchMentions   string `json:"launch_mentions"`
}

// UpdateSubscribedChannel takes a discord oauth token and information about a channel to change in the database.
func (a Api) UpdateSubscribedChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	token := r.Header.Get("Discord-Bearer-Token")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "no Discord-Bearer-Token header"})
		return
	}

	var requestedUpdate updateSubscribedChannelJson
	err := json.NewDecoder(r.Body).Decode(&requestedUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "failed to decode JSON body"})
		return
	}

	guilds, err := a.discordClient.GetGuildList(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "error getting information from Discord API"})
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
	//  provided guild) and then being able to edit a channel from another guild.
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

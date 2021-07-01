package api

import (
	"encoding/json"
	"net/http"
)

// updateChannelJson is a struct to marshal the api request data into.
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
		endWithResponse(w, responseBadJson)
		return
	}

	allowedToEdit := false
	for _, guild := range guilds {
		if guild.HasAdminPerms() && guild.ID == requestedUpdate.GuildID {
			allowedToEdit = true
		}
	}
	if !allowedToEdit {
		endWithResponse(w, responseNotAdmin)
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
		endWithResponse(w, responseDatabaseError)
		return
	}
	if !changed {
		endWithResponse(w, responseChannelNotInGuild)
		return
	}

	endWithResponse(w, responseAllOk)
}

package api

import (
	"encoding/json"
	"net/http"
)

// genericResponse is for any generic response we send from the API. All error messages should be user friendly.
// The default value for bool is false so if you want to return an error you don't need to worry about setting Success.
type genericResponse struct {
	Success    bool   `json:"success"`
	Error      string `json:"error,omitempty"`
	StatusCode int    `json:"status_code"` // Should be a HTTP code.
}

// subscribedResponse is the API response specifically for the subscribed channels API route.
type subscribedResponse struct {
	genericResponse
	Subscribed map[string]*guildDetails `json:"subscribed"`
}

// newSubscribedResponse initializes a new subscribedResponse.
func newSubscribedResponse() subscribedResponse {
	r := subscribedResponse{}
	r.Success = true
	r.Subscribed = make(map[string]*guildDetails)
	return r
}

// endWithResponse writes the response r to w.
// Doesn't actually "end" the ResponseWriter but it shouldn't be used after calling this.
func endWithResponse(w http.ResponseWriter, r interface{}) {
	// There must be a better way to do this.
	// TODO: Return json encode err?
	if resp, ok := r.(genericResponse); ok {
		if resp.StatusCode == 0 {
			resp.StatusCode = http.StatusOK
		}
		w.WriteHeader(resp.StatusCode)
		_ = json.NewEncoder(w).Encode(resp)
	} else if resp, ok := r.(subscribedResponse); ok {
		if resp.StatusCode == 0 {
			resp.StatusCode = http.StatusOK
		}
		w.WriteHeader(resp.StatusCode)
		_ = json.NewEncoder(w).Encode(resp)
	}
}

// Define some common responses.
var responseAllOk = genericResponse{Success: true}
var responseNoAuthHeader = genericResponse{Error: "no Authorization header", StatusCode: http.StatusUnauthorized}
var responseDiscordApiError = genericResponse{Error: "error getting information from Discord API: ", StatusCode: http.StatusServiceUnavailable}
var responseDatabaseError = genericResponse{Error: "database error :(", StatusCode: http.StatusInternalServerError}
var responseChannelNotInGuild = genericResponse{Error: "no channel with that ID in the given guild"}
var responseNotAdmin = genericResponse{Error: "you are not an admin in that server", StatusCode: http.StatusForbidden}
var responseNotAdminInAny = genericResponse{Error: "you do not have admin permissions in any guilds"}
var responseNoSubscribedInAny = genericResponse{Error: "you do not have any subscribed channels in guilds that you administrate"}
var responseBadJson = genericResponse{Error: "failed to decode JSON body", StatusCode: http.StatusBadRequest}

// Contains generic response related variables and functions.
// Other files in this directory may contain other response declarations.

package api

import (
	"encoding/json"
	"net/http"
)

// response defines the interface for API responses.
type response interface {
	finalize() int
}

// genericResponse is for any generic response we send from the API. All error messages should be user friendly.
// The default value for bool is false so if you want to return an error you don't need to worry about setting Success.
type genericResponse struct {
	Success    bool   `json:"success"`
	Error      string `json:"error,omitempty"`
	StatusCode int    `json:"status_code"` // Should be a HTTP code.
}

// finalize makes sure the genericResponse is ready to be sent to the client, and returns the used status code.
func (r *genericResponse) finalize() int {
	if r.StatusCode == 0 {
		r.StatusCode = http.StatusOK
	}
	return r.StatusCode
}

// endWithResponse writes the response r to w. The struct that implements response must be passed as a pointer.
// This is because *genericResponse implements finalize and genericResponse doesn't.
// This function doesn't actually "end" the ResponseWriter but it shouldn't be used after calling this.
func endWithResponse(w http.ResponseWriter, r response) {
	// TODO: Return json encode err?
	w.WriteHeader(r.finalize())
	_ = json.NewEncoder(w).Encode(r)
}

// Define some common responses.
var responseAllOk = &genericResponse{Success: true}
var responseNoAuthHeader = &genericResponse{Error: "No Authorization header", StatusCode: http.StatusUnauthorized}
var responseDiscordApiError = &genericResponse{Error: "Error getting information from Discord API: ", StatusCode: http.StatusServiceUnavailable}
var responseDatabaseError = &genericResponse{Error: "Database error :(", StatusCode: http.StatusInternalServerError}
var responseChannelNotInGuild = &genericResponse{Error: "No channel with that ID in the given guild"}
var responseNotAdmin = &genericResponse{Error: "You are not an admin in that server", StatusCode: http.StatusForbidden}
var responseNotAdminInAny = &genericResponse{Error: "You do not have admin permissions in any guilds"}
var responseNoSubscribedInAny = &genericResponse{Error: "You do not have any subscribed channels in guilds that you administrate"}
var responseBadJson = &genericResponse{Error: "Failed to decode JSON body", StatusCode: http.StatusBadRequest}

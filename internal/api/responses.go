// Contains response definitions and related variables and functions.

package api

import (
	"net/http"

	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/gin-gonic/gin"
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

// userInfoResponse is the API response for the userinfo API route.
type userInfoResponse struct {
	genericResponse
	UserInfo discord.UserInfo `json:"user_info"`
}

// subscribedResponse is the API response for the subscribed channels API route.
type subscribedResponse struct {
	genericResponse
	Subscribed map[string]*subscribedResponseGuildDetails `json:"subscribed"`
}

// subscribedResponseGuildDetails holds information about a guild for subscribedResponse.
type subscribedResponseGuildDetails struct {
	Name               string                             `json:"name"`
	Icon               string                             `json:"icon"`
	SubscribedChannels []database.SubscribedChannelRecord `json:"subscribed_channels"`
}

// statsResponse is the API response for the stats API route.
type statsResponse struct {
	genericResponse
	CountRecords []database.CountRecord `json:"counts"`
	ActionCounts []database.ActionCount `json:"action_counts"`
}

// finalize makes sure the genericResponse is ready to be sent to the client, and returns the used status code.
func (r *genericResponse) finalize() int {
	if r.StatusCode == 0 {
		r.StatusCode = http.StatusOK
	}
	return r.StatusCode
}

// endWithResponse writes the response r to c using AbortWithStatusJSON.
// The struct that implements response must be passed as a pointer, because *genericResponse implements finalize and
// genericResponse doesn't. c should not be used again after calling this function.
func endWithResponse(c *gin.Context, r response) {
	c.AbortWithStatusJSON(r.finalize(), r)
}

// Common responses.
var responseAllOk = &genericResponse{Success: true}
var responseNoSession = &genericResponse{Error: "No active login session found", StatusCode: http.StatusUnauthorized}
var responseDiscordApiError = &genericResponse{Error: "Error getting information from Discord API: ", StatusCode: http.StatusServiceUnavailable}
var responseDatabaseError = &genericResponse{Error: "Database error :(", StatusCode: http.StatusInternalServerError}
var responseChannelNotInGuild = &genericResponse{Error: "No channel with that ID in the given guild"}
var responseNotAdmin = &genericResponse{Error: "You are not an admin in that server", StatusCode: http.StatusForbidden}
var responseNotAdminInAny = &genericResponse{Error: "You do not have admin permissions in any guilds"}
var responseNoSubscribedInAny = &genericResponse{Error: "You do not have any subscribed channels in guilds that you administrate"}
var responseBadJson = &genericResponse{Error: "Failed to decode JSON body", StatusCode: http.StatusBadRequest}
var responseInternalError = &genericResponse{Error: "Internal server error", StatusCode: http.StatusInternalServerError}

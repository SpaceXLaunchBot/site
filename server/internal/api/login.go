package api

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Take the "token" parameter, use it to request guilds from Discord API, filter only for guilds that have
	//  subscribed channels and that the user is an admin in, then return data for frontend.
}

package api

import (
	"encoding/json"
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"net/http"
)

// TODO: Maybe cache in Api http handlers instead of in the discord package.

// apiResponse is for any generic response we send from the API. All error messages should be user friendly.
// The default value for bool is false so if you want to return an error you don't need to worry about setting Success.
type apiResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Api contains methods that interface with the database and the Discord API through a REST API.
type Api struct {
	db            database.Db
	discordClient discord.Client
}

// NewApi creates a new Api.
func NewApi(db database.Db, client discord.Client) Api {
	return Api{
		db:            db,
		discordClient: client,
	}
}

// getGuildList acts like a middleware and gets a GuildList using the authorization header (or sends an error to the client).
func (a Api) getGuildList(w http.ResponseWriter, r *http.Request) (list discord.GuildList, sentErr bool) {
	token := r.Header.Get("authorization")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "no authorization header"})
		return discord.GuildList{}, true
	}

	guilds, err := a.discordClient.GetGuildList(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apiResponse{Error: "error getting information from Discord API"})
		return discord.GuildList{}, true
	}

	return guilds, false
}

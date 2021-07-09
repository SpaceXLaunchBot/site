package api

import (
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"net/http"
)

// TODO: Maybe cache in Api http handlers instead of in the discord package.

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

// getGuildList acts like a middleware and gets a GuildList using the Authorization header (or sends an error to the client).
func (a Api) getGuildList(w http.ResponseWriter, r *http.Request) (list discord.GuildList, sentErr bool) {
	token := r.Header.Get("Authorization")
	if token == "" {
		endWithResponse(w, responseNoAuthHeader)
		return discord.GuildList{}, true
	}

	guilds, err := a.discordClient.GetGuildList(token)
	if err != nil {
		resp := responseDiscordApiError
		// Add context to general error message.
		resp.Error += err.Error()
		endWithResponse(w, resp)
		return discord.GuildList{}, true
	}

	return guilds, false
}

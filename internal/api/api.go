package api

import (
	"github.com/psidex/SpaceXLaunchBotSite/internal/database"
	"github.com/psidex/SpaceXLaunchBotSite/internal/discord"
)

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

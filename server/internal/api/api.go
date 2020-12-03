package api

import (
	"github.com/psidex/SpaceXLaunchBotSite/internal/database"
	"github.com/psidex/SpaceXLaunchBotSite/internal/discord"
)

// apiError is the response when an error happens in an API request. All error messages should be user friendly.
type apiError struct {
	Error string `json:"error"`
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

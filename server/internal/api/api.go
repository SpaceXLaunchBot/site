package api

import (
	"github.com/psidex/SpaceXLaunchBotSite/internal/database"
)

// apiError is the response when an error happens in an API request. All error messages should be user friendly.
type apiError struct {
	Error string `json:"error"`
}

// Api contains methods that interface with the database through a REST API.
type Api struct {
	db database.Db
}

// NewApi creates a new Api.
func NewApi(db database.Db) Api {
	return Api{
		db: db,
	}
}

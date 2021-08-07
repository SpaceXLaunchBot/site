package database

import (
	"fmt"

	"github.com/SpaceXLaunchBot/site/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib" // Provides Postgres "pgx" driver for sql
	"github.com/jmoiron/sqlx"          // Mainly used so we can marshal rows straight to structs
)

// TODO: Use transactions?
// TODO: Invalidate old sessions, maybe have a goroutine that runs every hour or something.

// Db is a wrapper around sqlx.DB which provides methods for interacting with the project specific tables.
type Db struct {
	sqlxHandle *sqlx.DB
}

// NewDb creates a new Db.
func NewDb(c config.Config) (Db, error) {
	conStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		c.DbUser, c.DbPass, c.DbHost, c.DbPort, c.DbName,
	)
	sqlxHandle, err := sqlx.Connect("pgx", conStr)
	if err != nil {
		return Db{}, err
	}
	return Db{sqlxHandle}, nil
}

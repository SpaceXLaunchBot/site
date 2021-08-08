package database

import (
	"fmt"
	"time"

	"github.com/SpaceXLaunchBot/site/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib" // Provides Postgres "pgx" driver for sql
	"github.com/jmoiron/sqlx"          // Mainly used so we can marshal rows straight to structs
)

// Db is a wrapper around sqlx.DB which provides methods for interacting with SLB specific tables.
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
	d := Db{sqlxHandle}
	go d.sessionReaper(time.Hour)
	return d, nil
}

// sessionReaper checks every dur for sessions that are more than 6 months old and rms them.
func (d Db) sessionReaper(dur time.Duration) {
	for {
		query := `delete from sessions where session_creation_time < now() - INTERVAL '1 MINUTE';`

		res, _ := d.sqlxHandle.Exec(query)
		// TODO: Logging for errors.
		// if err != nil {

		_, _ = res.RowsAffected()
		// if err != nil {

		time.Sleep(dur)
	}
}

package database

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/psidex/SpaceXLaunchBotSite/internal/config"
)

// SubscribedChannel represents a subscribed channel from the database. Can be marshaled to json (GuildId won't be).
type SubscribedChannel struct {
	Id               string         `db:"channel_id" json:"id"`
	GuildId          string         `db:"guild_id" json:"-"`
	Name             string         `db:"channel_name" json:"name"`
	NotificationType string         `db:"notification_type" json:"notification_type"`
	LaunchMentions   sql.NullString `db:"launch_mentions" json:"launch_mentions"`
}

// Db represents a connection to the database and provides methods for interacting with it.
type Db struct {
	db *sqlx.DB
}

// NewDb creates a new Db.
func NewDb(c config.Config) (Db, error) {
	conStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", c.DbUser, c.DbPass, c.DbHost, c.DbName)
	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		return Db{}, err
	}
	return Db{db}, nil
}

// SubscribedChannels returns a slice of SubscribedChannels that are exist in the given guild ids.
func (d Db) SubscribedChannels(guildIds []string) ([]SubscribedChannel, error) {
	var channels []SubscribedChannel
	query, args, err := sqlx.In(`
		SELECT *
		FROM subscribed_channels
		WHERE guild_id in (?);`,
		guildIds,
	)
	if err != nil {
		return channels, err
	}

	// Rebind takes the general form of query created by In and converts to what Postgres wants.
	query = d.db.Rebind(query)

	err = d.db.Select(&channels, query, args...)
	if err != nil {
		return channels, err
	}
	return channels, nil
}

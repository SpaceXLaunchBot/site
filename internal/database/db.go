package database

import (
	"database/sql"
	"fmt"
	"github.com/SpaceXLaunchBot/site/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib" // Provides Postgres "pgx" driver for sql
	"github.com/jmoiron/sqlx"          // Mainly used so we can marshal rows straight to structs
	"strings"
)

// SubscribedChannel represents a subscribed channel from the database.
type SubscribedChannel struct {
	Id               string         `db:"channel_id"`
	GuildId          string         `db:"guild_id"`
	Name             string         `db:"channel_name"`
	NotificationType string         `db:"notification_type"`
	LaunchMentions   sql.NullString `db:"launch_mentions"`
}

// Db represents a connection to the database and provides methods for interacting with it.
type Db struct {
	db *sqlx.DB
}

// NewDb creates a new Db.
func NewDb(c config.Config) (Db, error) {
	conStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		c.DbUser, c.DbPass, c.DbHost, c.DbPort, c.DbName,
	)
	db, err := sqlx.Connect("pgx", conStr)
	if err != nil {
		return Db{}, err
	}
	return Db{db}, nil
}

// SubscribedChannels returns a slice of SubscribedChannels that exist in the given guild ids.
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

	// Rebind takes query with sql bind vars ("?") from sqlx.In and turns them into $1, $2, etc. for Postgres.
	query = d.db.Rebind(query)

	err = d.db.Select(&channels, query, args...)
	if err != nil {
		return channels, err
	}
	return channels, nil
}

// UpdateSubscribedChannel sets the notification type and launch mentions for a given channel ID.
// guildId is required to ensure that the channel exists in that guild.
func (d Db) UpdateSubscribedChannel(channelId, guildId, notificationType, launchMentions string) (changed bool, err error) {
	sqlLaunchMentions := sql.NullString{
		String: launchMentions,
		Valid:  strings.TrimSpace(launchMentions) != "",
	}
	const query = `
		UPDATE subscribed_channels SET (notification_type, launch_mentions) = ($1, $2)
		WHERE channel_id = $3 AND guild_id = $4;`
	res, err := d.db.Exec(query, notificationType, sqlLaunchMentions, channelId, guildId)
	if err != nil {
		return false, err
	}
	num, err := res.RowsAffected()
	return num > 0, err
}

// DeleteSubscribedChannel removes a given channel from the subscribed_channels table.
func (d Db) DeleteSubscribedChannel(channelId, guildId string) (changed bool, err error) {
	query := "DELETE FROM subscribed_channels WHERE channel_id = $1 AND guild_id = $2;"
	res, err := d.db.Exec(query, channelId, guildId)
	if err != nil {
		return false, err
	}
	num, err := res.RowsAffected()
	return num > 0, err
}

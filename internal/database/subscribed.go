package database

import "github.com/jmoiron/sqlx"

// SubscribedChannelRecord represents a record of a subscribed channel from the database.
// Can be marshalled straight to JSON, be careful of nil pointers.
type SubscribedChannelRecord struct {
	Id               string  `db:"channel_id" json:"id"`
	GuildId          string  `db:"guild_id"`
	Name             string  `db:"channel_name" json:"name"`
	NotificationType string  `db:"notification_type" json:"notification_type"`
	LaunchMentions   *string `db:"launch_mentions" json:"launch_mentions"` // Pointer because it can be NULL in the db.
}

// SubscribedChannels returns a slice of SubscribedChannels that exist in the given guild ids.
func (d Db) SubscribedChannels(guildIds []string) ([]SubscribedChannelRecord, error) {
	var channels []SubscribedChannelRecord
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
	query = d.sqlxHandle.Rebind(query)

	err = d.sqlxHandle.Select(&channels, query, args...)
	if err != nil {
		return channels, err
	}
	return channels, nil
}

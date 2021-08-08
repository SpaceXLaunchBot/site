package database

import "github.com/jmoiron/sqlx"

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

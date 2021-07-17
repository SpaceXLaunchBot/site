package database

import (
	"database/sql"
	"strings"
)

// UpdateSubscribedChannel sets the notification type and launch mentions for a given channel ID.
// guildId is required to ensure that the channel exists in that guild.
func (d Db) UpdateSubscribedChannel(channelId, guildId, notificationType, launchMentions string) (changed bool, err error) {
	// Set it to NULL in the db if it doesn't exist.
	// TODO: Don't use NullString if we are using sqlx.In?
	sqlLaunchMentions := sql.NullString{
		String: launchMentions,
		Valid:  strings.TrimSpace(launchMentions) != "",
	}
	// TODO: Use ? and then rebind.
	const query = `
		UPDATE subscribed_channels SET (notification_type, launch_mentions) = ($1, $2)
		WHERE channel_id = $3 AND guild_id = $4;`
	res, err := d.sqlxHandle.Exec(query, notificationType, sqlLaunchMentions, channelId, guildId)
	if err != nil {
		return false, err
	}
	num, err := res.RowsAffected()
	return num > 0, err
}

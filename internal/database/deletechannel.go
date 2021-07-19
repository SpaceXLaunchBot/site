package database

// DeleteSubscribedChannel removes a given channel from the subscribed_channels table.
func (d Db) DeleteSubscribedChannel(channelId, guildId string) (changed bool, err error) {
	query := "DELETE FROM subscribed_channels WHERE channel_id = ? AND guild_id = ?;"
	query = d.sqlxHandle.Rebind(query)
	res, err := d.sqlxHandle.Exec(query, channelId, guildId)
	if err != nil {
		return false, err
	}
	num, err := res.RowsAffected()
	return num > 0, err
}

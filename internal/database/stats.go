package database

// CountRecord represents a record of guild / subscribed channel counts from the database.
// Can be marshalled straight to JSON.
type CountRecord struct {
	GuildCount      int    `db:"guild_count" json:"g"`
	SubscribedCount int    `db:"subscribed_count" json:"s"`
	Date            string `db:"date" json:"d"`
}

// Stats WIP.
func (d Db) Stats() ([]CountRecord, error) {
	var ms []CountRecord

	err := d.sqlxHandle.Select(&ms, `
		SELECT
			guild_count, subscribed_count,
			to_char("time", 'YYYY-MM-DD HH:00:00') as "date"
		FROM counts;`,
	)

	if err != nil {
		return ms, err
	}
	return ms, nil
}

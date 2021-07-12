package database

// CountRecord represents a record of guild / subscribed channel counts from the database.
// Can be marshalled straight to JSON.
type CountRecord struct {
	GuildCount      int    `db:"guild_count" json:"guild_count"`
	SubscribedCount int    `db:"subscribed_count" json:"subscribed_count"`
	Date            string `db:"date" json:"date"`
	Hour            int    `db:"hour" json:"hour"`
}

// Stats WIP.
func (d Db) Stats() ([]CountRecord, error) {
	var ms []CountRecord

	err := d.sqlxHandle.Select(&ms, `
		SELECT
			guild_count, subscribed_count,
			to_char("time", 'YYYY-MM-DD') as "date",
			EXTRACT(HOUR FROM "time") as hour
		FROM counts;`,
	)

	if err != nil {
		return ms, err
	}
	return ms, nil
}

package database

// Stats WIP.
func (d Db) Stats() ([]CountRecord, []ActionCount, error) {
	var counts []CountRecord
	var actionCounts []ActionCount

	err := d.sqlxHandle.Select(&counts, `
		SELECT
			guild_count, subscribed_count,
			to_char("time", 'YYYY-MM-DD HH24:00:00') AS "date"
		FROM counts;`,
	)
	if err != nil {
		return counts, actionCounts, err
	}

	// Currently we just select commands used, maybe rename from ActionCount to CommandCount (and others)?
	err = d.sqlxHandle.Select(&actionCounts, `
		SELECT
			replace(replace(replace(action, 'command_', ''), '_cmd', ''), '_', '') as "action_formatted",
			count(action) as "count"
		FROM metrics
		WHERE action like 'command_%'
		GROUP BY "action_formatted";`,
	)
	if err != nil {
		return counts, actionCounts, err
	}

	return counts, actionCounts, nil
}

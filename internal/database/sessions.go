package database

import "time"

// SetSession creates a session record in the database.
func (d Db) SetSession(id string, accessTokenEncrypted []byte, expiresIn int, refreshTokenEncrypted []byte) (changed bool, err error) {
	expiresAt := time.Unix(time.Now().Unix()+int64(expiresIn), 0)
	query := `
		INSERT INTO sessions
		    (session_id, access_token_encrypted, access_token_expires_at, refresh_token_encrypted)
		VALUES
		    (?, ?, ?, ?);`

	query = d.sqlxHandle.Rebind(query)
	res, err := d.sqlxHandle.Exec(
		query,
		id,
		accessTokenEncrypted,
		expiresAt,
		refreshTokenEncrypted,
	)
	if err != nil {
		return false, err
	}

	num, err := res.RowsAffected()
	return num > 0, err
}

// GetSession gets a session record from the database with the given session id.
func (d Db) GetSession(id string) (exists bool, record SessionRecord, err error) {
	var sessionRecords []SessionRecord
	var session SessionRecord

	query := `SELECT * FROM sessions WHERE session_id=?;`

	query = d.sqlxHandle.Rebind(query)
	err = d.sqlxHandle.Select(&sessionRecords, query, id)
	if err != nil {
		return false, session, err
	}
	if len(sessionRecords) == 0 {
		return false, session, nil
	}

	session = sessionRecords[0]
	if session.SessionId == "" {
		// Not sure if this is actually something that would ever happen.
		return false, session, nil
	}

	return true, session, nil
}

// RemoveSession remove a session record from the database with the given session id.
func (d Db) RemoveSession(id string) (deleted bool, err error) {
	query := `DELETE FROM sessions WHERE session_id=?;`

	query = d.sqlxHandle.Rebind(query)
	res, err := d.sqlxHandle.Exec(query, id)
	if err != nil {
		return false, err
	}

	num, err := res.RowsAffected()
	return num > 0, err
}

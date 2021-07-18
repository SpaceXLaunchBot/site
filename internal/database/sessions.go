package database

import (
	"github.com/SpaceXLaunchBot/site/internal/encryption"
	"github.com/jmoiron/sqlx"
	"time"
)

// SessionRecord represents a record in the sessions table.
type SessionRecord struct {
	Session               string `db:"session"` // The uuid string that is in the clients cookie.
	AccessToken           string
	AccessTokenEncrypted  []byte    `db:"access_token_encrypted"`
	AccessTokenExpiresAt  time.Time `db:"access_token_expires_at"`
	RefreshToken          string
	RefreshTokenEncrypted []byte    `db:"refresh_token_encrypted"`
	CreationTime          time.Time `db:"creation_time"`
}

// SetSession creates a session record in the database.
func (d Db) SetSession(id string, key []byte, accessToken string, expiresIn int, refreshToken string) (changed bool, err error) {
	accessTokenEncrypted, err := encryption.Encrypt(key, []byte(accessToken))
	if err != nil {
		return false, err
	}

	refreshTokenEncrypted, err := encryption.Encrypt(key, []byte(refreshToken))
	if err != nil {
		return false, err
	}

	expiresAt := time.Unix(time.Now().Unix()+int64(expiresIn), 0)
	query := `
		INSERT INTO sessions
		    ("session", access_token_encrypted, access_token_expires_at, refresh_token_encrypted)
		VALUES
		    (?, ?, ?, ?);
	`

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
func (d Db) GetSession(sessionId string, key []byte) (exists bool, record SessionRecord, err error) {
	var sessionRecords []SessionRecord
	var session SessionRecord

	query, args, err := sqlx.In(`SELECT * FROM sessions WHERE "session"=(?);`, sessionId)
	if err != nil {
		return false, session, err
	}

	query = d.sqlxHandle.Rebind(query)
	err = d.sqlxHandle.Select(&sessionRecords, query, args...)
	if err != nil {
		return false, session, err
	}
	if len(sessionRecords) == 0 {
		return false, session, nil
	}

	session = sessionRecords[0]
	if session.Session == "" {
		// Not sure if this is actually something that would ever happen.
		// We return nil in these because the error is not a database error, and we check for exists anyway.
		return false, session, nil
	}

	if session.AccessTokenExpiresAt.After(time.Now()) == false {
		// Everything is valid but our access token is expired.
		// TODO: Attempt to refresh with refresh token. Not sure where in codebase to do this.
		_, err = d.RemoveSession(sessionId)
		return false, session, nil
	}

	accessTokenBytes, err := encryption.Decrypt(key, session.AccessTokenEncrypted)
	if err != nil {
		return false, session, nil
	}

	refreshTokenBytes, err := encryption.Decrypt(key, session.RefreshTokenEncrypted)
	if err != nil {
		return false, session, nil
	}

	session.AccessToken = string(accessTokenBytes)
	session.RefreshToken = string(refreshTokenBytes)

	return true, session, nil
}

// RemoveSession remove a session record from the database with the given session id.
func (d Db) RemoveSession(sessionId string) (deleted bool, err error) {
	query, args, err := sqlx.In(`DELETE FROM sessions WHERE "session"=(?);`, sessionId)
	if err != nil {
		return false, err
	}

	query = d.sqlxHandle.Rebind(query)
	res, err := d.sqlxHandle.Exec(query, args...)
	if err != nil {
		return false, err
	}

	num, err := res.RowsAffected()
	return num > 0, err
}

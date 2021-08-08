// Contains database record definitions and related variables and functions.

package database

import "time"

// SubscribedChannelRecord represents a record in the subscribed channel table.
// Can be marshalled straight to JSON, be careful of nil pointers.
type SubscribedChannelRecord struct {
	Id               string  `db:"channel_id" json:"id"`
	GuildId          string  `db:"guild_id"`
	Name             string  `db:"channel_name" json:"name"`
	NotificationType string  `db:"notification_type" json:"notification_type"`
	LaunchMentions   *string `db:"launch_mentions" json:"launch_mentions"` // Pointer because it can be NULL in the db.
}

// SessionRecord represents a record in the sessions table.
// The 2 non-marshalled-to fields are because we pass this around and we will need store the unencrypted values.
type SessionRecord struct {
	SessionId             string    `db:"session_id"`
	SessionCreationTime   time.Time `db:"session_creation_time"`
	AccessToken           string
	AccessTokenEncrypted  []byte    `db:"access_token_encrypted"`
	AccessTokenExpiresAt  time.Time `db:"access_token_expires_at"`
	RefreshToken          string
	RefreshTokenEncrypted []byte    `db:"refresh_token_encrypted"`
	RefreshTime           time.Time `db:"refresh_time"`
}

// CountRecord represents a record in the counts table.
// Can be marshalled straight to JSON.
type CountRecord struct {
	GuildCount      int    `db:"guild_count" json:"g"`
	SubscribedCount int    `db:"subscribed_count" json:"s"`
	Date            string `db:"date" json:"d"`
}

// ActionCount represents a formatted record from the metrics table.
// Can be marshalled straight to JSON.
type ActionCount struct {
	Action string `db:"action_formatted" json:"a"`
	Count  int    `db:"count" json:"c"`
}

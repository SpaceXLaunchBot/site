// Contains JSON request definitions and related variables and functions.

package api

// deleteChannelRequest is a struct to marshal the api request data into.
type deleteChannelRequest struct {
	ID      string `json:"id"`
	GuildID string `json:"guild_id"`
}

// updateChannelRequest is a struct to marshal the api request data into.
type updateChannelRequest struct {
	ID               string `json:"id"`
	GuildID          string `json:"guild_id"`
	NotificationType string `json:"notification_type"`
	LaunchMentions   string `json:"launch_mentions"`
}

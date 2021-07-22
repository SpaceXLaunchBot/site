package discord

import (
	"encoding/json"
)

// Guild represents information about a DiscordClient guild that a user is in.
type Guild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Permissions int    `json:"permissions"`
	//Owner          bool     `json:"owner"`
	//Features       []string `json:"features"`
	//PermissionsNew string   `json:"permissions_new"`
}

// HasAdminPerms determines if the user has admin permissions in the guild.
func (g Guild) HasAdminPerms() bool {
	// Admin is 0x00000008: https://discord.com/developers/docs/topics/permissions
	return 8&g.Permissions != 0
}

// GuildList represents a list of Guilds.
type GuildList []Guild

// GetGuildList returns a GuildList of guilds that the user of the token is in.
func (c Client) GetGuildList(bearerToken string) (GuildList, error) {
	endpoint := "/users/@me/guilds"
	body, err := c.apiRequestGet(endpoint, bearerToken)
	if err != nil {
		return GuildList{}, err
	}

	guildList := GuildList{}
	if err = json.Unmarshal(body, &guildList); err != nil {
		return GuildList{}, err
	}

	return guildList, nil
}

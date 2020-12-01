package discord

import (
	"encoding/json"
)

// Guild represents information about a Discord guild.
type Guild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Permissions int    `json:"permissions"`
	//Owner          bool     `json:"owner"`
	//Features       []string `json:"features"`
	//PermissionsNew string   `json:"permissions_new"`
}

// GuildList represents a list of Guilds.
type GuildList []Guild

// GetGuildList returns a GuildList of guilds that the user of the token is in.
func GetGuildList(bearerToken string) (GuildList, error) {
	endpoint := "/users/@me/guilds"
	body, err := apiRequest(endpoint, bearerToken)
	if err != nil {
		return GuildList{}, err
	}

	guildList := GuildList{}
	if err = json.Unmarshal(body, &guildList); err != nil {
		return GuildList{}, err
	}
	return guildList, nil
}

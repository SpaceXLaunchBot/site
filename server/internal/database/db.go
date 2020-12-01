package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/psidex/SpaceXLaunchBotSite/internal/config"
)

const channelsFromGuildSql = `
SELECT s.channel_name,
	   s.channel_id,
	   s.notification_type,
       s.launch_mentions
FROM guild AS g
JOIN subscribed_channels AS s ON g.guild_id = s.guild_id
WHERE g.guild_id = $1;`

const guildNameSql = "SELECT guild_name FROM guild where guild_id = $1;"

// Channel represents a subscribed channel from the database, can be marshaled to json.
type Channel struct {
	Name             string `json:"name"`
	Id               string `json:"id"`
	NotificationType string `json:"notification_type"`
	LaunchMentions   string `json:"launch_mentions"`
}

// Db represents a connection to the database and provides methods for interacting with it.
type Db struct {
	connString string
}

// NewDb creates a new Db.
func NewDb(c config.Config) Db {
	// TODO: Maybe use connection pooling?
	return Db{fmt.Sprintf("postgresql://%s:%s@%s/%s", c.DbUser, c.DbPass, c.DbHost, c.DbName)}
}

// GuildName returns the name of a guild from the database given the id.
func (d Db) GuildName(guildId string) (string, error) {
	conn, err := pgx.Connect(context.Background(), d.connString)
	if err != nil {
		return "", err
	}
	defer conn.Close(context.Background())

	var name string
	err = conn.QueryRow(context.Background(), guildNameSql, guildId).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

// SubscribedChannels returns the information for channels subscribed in the given guild id.
func (d Db) SubscribedChannels(guildId string) ([]Channel, error) {
	var channels []Channel

	conn, err := pgx.Connect(context.Background(), d.connString)
	if err != nil {
		return channels, err
	}
	defer conn.Close(context.Background())

	channelRows, err := conn.Query(context.Background(), channelsFromGuildSql, guildId)
	if err != nil {
		return channels, err
	}
	defer channelRows.Close()

	for channelRows.Next() {
		var n, id, nt, lm string
		_ = channelRows.Scan(&n, &id, &nt, &lm)
		channels = append(channels, Channel{
			Name:             n,
			Id:               id,
			NotificationType: nt,
			LaunchMentions:   lm,
		})
	}

	return channels, nil
}

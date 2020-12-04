# API Docs

All requests require the header "Discord-Bearer-Token".

## Error

This will be returned from any endpoint if an error occurs.

### Response

```json
{
  "success": false,
  "error": "Error message"
}
```

## `GET` `/api/guildswithsubscribed`

### Response

```json
{
    "{discord snowflake}": {
        "name": "Guild Name",
        "icon": "https://cdn.discordapp.com/icons/{discord snowflake}/{discord snowflake}.png",
        "subscribed_channels": [
            {
            "id": "{discord snowflake}",
            "name": "Channel Name",
            "notification_type": "{all|schedule|launch}",
            "launch_mentions": "mentions"
            },
            {
              etc.
            }
        ]
    },
    {
      etc.
    }
}
```

## `POST` `/api/updatesubscribedchannel`

### Request Body

```json
{
	"id": "{discord snowflake}",
	"guild_id": "{discord snowflake}",
	"notification_type": "{all|schedule|launch}",
	"launch_mentions": "mentions"
}
```

### Response

```json
{
  "success": true,
  "error": ""
}
```

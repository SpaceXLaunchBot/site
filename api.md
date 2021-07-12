# API Docs

All requests to the API require the header `"authorization"` to be set with a valid Discord OAuth token. The token
should be authorized with the `identity` and `guilds` scopes. This token is used to verify that you are an admin in any
given channels/servers that you want to view/update/delete.

## Error

If an error ever occurs with any endpoint, the response will take this form.

`status_code` will be the HTTP code that was returned by the server.

```json
{
  "success": false,
  "error": "User friendly error message",
  "status_code": 401
}
```

## `GET` `/api/subscribed`

Returns information about guilds and channels that the user is in that are subscribed to the notification service.

### Response

This response can contain multiple guilds with multiple subscribed channels.

```json
{
  "success": true,
  "status_code": 200,
  "subscribed": {
    "{discord snowflake}": {
      "name": "Guild Name",
      "icon": "https://cdn.discordapp.com/icons/{discord snowflake}/{discord snowflake}.png",
      "subscribed_channels": [
        {
          "id": "{discord snowflake}",
          "name": "Channel Name",
          "notification_type": "{all|schedule|launch}",
          "launch_mentions": "mentions"
        }
      ]
    },
    "etc": {...}
  }
}
```

## `DELETE` `/api/channel`

Removes a subscribed channel from the database that matches the given information.

### Request Body

```json
{
	"id": "{discord snowflake}",
	"guild_id": "{discord snowflake}"
}
```

### Response

```json
{
  "success": true,
  "status_code": 200
}
```

## `PUT` `/api/channel`

Updates a subscribed channel in the database with the given information.

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
  "status_code": 200
}
```

## `GET` `/api/stats`

Returns some basic stats about the bot. Currently a WIP.

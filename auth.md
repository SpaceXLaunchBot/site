# Authentication

This document details the structure of the authentication system and how this website interfaces with Discords OAuth API.

This is mainly so I can read this in the future and understand what I've done, but it may be useful for other people as well :)

See [this discussion](https://security.stackexchange.com/a/77316) as to why I chose to do everything the way I did.

Error handling is assumed in the algorithms/examples below.

## Logging In

- User logs in to discord, returns to `/login` with a `code` query param
- Frontend route at `/login` extracts this param and sends it in a request to `/api/auth/login` 
- Backend receives this `code` and uses it to request OAuth tokens from Discord
- Backend receives access and refresh tokens
- A session ID is created
- A 32 byte encryption key is created
- The access and refresh tokens are encrypted with this key
- The session ID along with the encrypted tokens and metadata are inserted into the database
- The session ID and encryption key and given to the frontend in http-only cookies called `"session"` and `"key"`
- An "all ok" API response is given to the frontend by the backend
  
## Accessing session endpoints

- Frontend requests `/api/userinfo` (for example)
- Session middleware finds session id and key cookies
- The backend passes these to a database handler which retrieves and decrypts the frontend users session information (most importantly OAuth tokens)
- These tokens can now be used by the backend to request information from the Discord API about the user
- Backend responds with relevant data

## Logging out

- Frontend requests `/api/auth/logout`
- Session middleware finds session ID
- The session with the given ID is removed from the database
- The session id and key cookies are set to have empty values and expire immediately
- An "all ok" API response is given to the frontend by the backend

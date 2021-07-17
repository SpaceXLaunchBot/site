package api

import (
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

// TODO: Maybe cache in Api http handlers instead of in the discord package, reduces overall computation because we wont
//  have to go thru all the guild admin testing.

// Api contains methods that interface with the database and the Discord API through a REST API.
type Api struct {
	db            database.Db
	cache         *cache.Cache
	discordClient discord.Client
	// Host URI info for cookies.
	hostName string
	isHTTPS  bool
}

// NewApi creates a new Api.
func NewApi(db database.Db, client discord.Client, hostName, protocol string) Api {
	secure := true
	if protocol == "http:" {
		secure = false
	}
	return Api{
		db:            db,
		discordClient: client,
		cache:         cache.New(10*time.Second, 10*time.Minute),
		hostName:      hostName,
		isHTTPS:       secure,
	}
}

func (a Api) getSessionFromCookie(r *http.Request) (sessionExists bool, record database.SessionRecord, err error) {
	sessionCookie, err := r.Cookie("session")
	if err != nil {
		return false, record, err
	}
	return a.db.GetSession(sessionCookie.Value)
}

// getGuildList acts like a middleware and gets a GuildList using the Authorization header (or sends an error to the client).
func (a Api) getGuildList(w http.ResponseWriter, r *http.Request) (list discord.GuildList, sentErr bool) {
	exists, session, err := a.getSessionFromCookie(r)
	if err != nil || !exists {
		endWithResponse(w, responseNoSession)
		return discord.GuildList{}, true
	}

	// TODO: What happens if we have access token but it is expired or invalid?
	//  Same goes for in userinfo
	token := session.AccessToken
	guilds, err := a.discordClient.GetGuildList(token)
	if err != nil {
		resp := responseDiscordApiError
		// Add context to general error message.
		resp.Error += err.Error()
		endWithResponse(w, resp)
		return discord.GuildList{}, true
	}

	return guilds, false
}

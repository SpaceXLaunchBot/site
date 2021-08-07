package api

import (
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/gin-gonic/gin"
)

// UserInfo is the endpoint for getting the users username and avatar URL.
func (a Api) UserInfo(c *gin.Context) {
	session := c.MustGet("session").(database.SessionRecord)

	userInfo, err := a.discordClient.GetUserInfo(session.AccessToken)
	if err != nil {
		if err == discord.ErrBadAuth {
			a.endWithInvalidateSession(c, session.SessionId)
			return
		}
		resp := responseDiscordApiError
		resp.Error += err.Error()
		endWithResponse(c, resp)
		return
	}

	resp := &userInfoResponse{UserInfo: userInfo}
	resp.Success = true
	endWithResponse(c, resp)
}

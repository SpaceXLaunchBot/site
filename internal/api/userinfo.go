package api

import (
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"net/http"
)

type userInfoResponse struct {
	genericResponse
	UserInfo discord.UserInfo `json:"user_info"`
}

func (a Api) UserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	exists, session, err := a.getSessionFromCookie(r)
	if err != nil || !exists {
		endWithResponse(w, responseNoSession)
		return
	}

	userInfo, err := a.discordClient.GetUserInfo(session.AccessToken)
	if err != nil {
		resp := responseDiscordApiError
		resp.Error += err.Error()
		endWithResponse(w, resp)
		return
	}

	response := &userInfoResponse{UserInfo: userInfo}
	response.Success = true
	endWithResponse(w, response)
}

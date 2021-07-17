package discord

import (
	"encoding/json"
	"fmt"
)

type meRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	//Discriminator string `json:"discriminator"`
	//PublicFlags   int    `json:"public_flags"`
	//Flags         int    `json:"flags"`
	//Locale        string `json:"locale"`
	//MfaEnabled    bool   `json:"mfa_enabled"`
}

// UserInfo represents information the frontend needs about the logged in user.
type UserInfo struct {
	UserName  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

// GetUserInfo gets a UserInfo struct using the given bearerToken.
func (c Client) GetUserInfo(bearerToken string) (UserInfo, error) {
	var userInfoRequest meRequest
	var userInfo UserInfo

	endpoint := "/users/@me"

	body, err := c.apiRequestWithToken(endpoint, bearerToken)
	if err != nil {
		return userInfo, err
	}

	err = json.Unmarshal(body, &userInfoRequest)
	if err != nil {
		return userInfo, err
	}

	userInfo.UserName = userInfoRequest.Username
	userInfo.AvatarUrl = fmt.Sprintf(
		"https://cdn.discordapp.com/avatars/%s/%s.png", userInfoRequest.ID, userInfoRequest.Avatar,
	)
	return userInfo, nil
}

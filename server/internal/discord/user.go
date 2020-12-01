package discord

import (
	"encoding/json"
)

// User represents information about a Discord user.
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	//Discriminator string `json:"discriminator"`
	//PublicFlags   int    `json:"public_flags"`
	//Flags         int    `json:"flags"`
	//Locale        string `json:"locale"`
	//MfaEnabled    bool   `json:"mfa_enabled"`
}

// GetUser returns a User fpr the given token.
func GetUser(bearerToken string) (User, error) {
	endpoint := "/users/@me"
	body, err := apiRequest(endpoint, bearerToken)
	if err != nil {
		return User{}, err
	}

	user := User{}
	if err = json.Unmarshal(body, &user); err != nil {
		return User{}, err
	}
	return user, nil
}

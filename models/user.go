package models

import (
	"golara/app/database"
)

type User struct {
	ID          uint   `json:"id"`
	DisplayName string `json:"display_name"`
}

func (u *User) GetUserByToken(token string) error {
	var userToken UserToken
	err := database.DB().Preload("User").Where("access_token = ?", token).First(&userToken).Error
	if err != nil {
		return err
	}
	*u = *(&userToken.User)

	return err
}

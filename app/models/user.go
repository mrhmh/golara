package models

import "golara/core/db"

type User struct {
	ID          uint   `json:"id"`
	DisplayName string `json:"display_name"`
}

func (u *User) GetUserByToken(token string) error {
	var userToken UserToken
	err := db.DB().Preload("User").Where("access_token = ?", token).First(&userToken).Error
	if err != nil {
		return err
	}
	*u = *(&userToken.User)

	return err
}

package models

type UserToken struct {
	ID          uint   `json:"id"`
	UserId      int    `json:"user_id"`
	AccessToken string `json:"access_token"`
	User        User   `gorm:"foreignKey:UserId"`
}

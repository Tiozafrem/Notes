package model

import "time"

type DeviceUser struct {
	Id           int       `json:"-" db:"id"`
	Name         string    `json:"name_device" db:"name"`
	UserId       int       `json:"-" db:"user_id"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	Description  string    `json:"description" db:"description"`
	Expire       time.Time `json:"-" db:"expire"`
}

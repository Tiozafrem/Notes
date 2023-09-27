package model

// Table worh UserId and ListID
type UserList struct {
	Id     int `json:"id" db:"id"`
	UserId int `json:"user_id" db:"user_id"`
	ListId int `json:"list_id" db:"list_id"`
}

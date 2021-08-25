package model

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"requierd"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

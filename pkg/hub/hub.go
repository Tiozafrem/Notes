package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

type HubNotify interface {
	RegisterUser(user *User) error
	UnregisterUser(user *User) error
	EmitUser(userId int, payload string) error
	EmitUsers(usersId []int, payload string) error
}

type Hub struct {
	users map[*User]bool
}

type User struct {
	connection *websocket.Conn
	send       chan string
	userId     int
}

func NewUser(connection *websocket.Conn, userId int) *User {
	return &User{
		connection: connection,
		send:       make(chan string),
		userId:     userId,
	}
}

func NewHub() *Hub {
	return &Hub{
		users: make(map[*User]bool),
	}
}

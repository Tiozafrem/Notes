package hub

import (
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/slices"
)

func (h *Hub) RegisterUser(user *User) error {
	h.users[user] = true
	go user.write()
	return nil
}

func (h *Hub) UnregisterUser(user *User) error {
	_, ok := h.users[user]
	if ok {
		delete(h.users, user)
		close(user.send)
	}
	return nil
}

func (h *Hub) EmitUser(userId int, payload string) error {
	for user := range h.users {
		if user.userId == userId {
			select {
			case user.send <- payload:
			default:
				close(user.send)
				delete(h.users, user)
			}
		}
	}
	return nil
}

func (h *Hub) EmitUsers(userId []int, payload string) error {
	for user := range h.users {
		if slices.Contains(userId, user.userId) {
			select {
			case user.send <- payload:
			default:
				close(user.send)
				delete(h.users, user)
			}
		}
	}
	return nil
}

func (u *User) write() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		u.connection.Close()
		u.connection.Close()
	}()

	for {
		select {
		case payload, ok := <-u.send:
			u.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				u.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := u.connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write([]byte(payload))
			n := len(u.send)
			for i := 0; i < n; i++ {
				w.Write([]byte(payload))
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			u.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

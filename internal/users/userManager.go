package users

import "github.com/gorilla/websocket"

type UserManager struct {
	Users map[*websocket.Conn]bool
}

func NewUserManager() *UserManager {
	return &UserManager{Users: make(map[*websocket.Conn]bool)}
}

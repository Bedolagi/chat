package server

import (
	"chat/internal/users"
	"net/http"
	"sync"
)

type ServerManager struct {
	usersManager users.UserManager
	broadcast    chan string
	mutex        sync.Mutex
}

func NewServerManager() *ServerManager {
	return &ServerManager{usersManager: *users.NewUserManager(),
		broadcast: make(chan string),
	}
}

func (serverManager *ServerManager) Start() {
	http.HandleFunc("/ws", serverManager.handleConnections)

	go serverManager.handleMessages()

	http.ListenAndServe(":8080", nil)
}

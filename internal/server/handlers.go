package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func (serverManager *ServerManager) handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	serverManager.mutex.Lock()
	serverManager.usersManager.Users[ws] = true
	serverManager.mutex.Unlock()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			serverManager.mutex.Lock()
			delete(serverManager.usersManager.Users, ws)
			serverManager.mutex.Unlock()
			break
		}
		serverManager.broadcast <- string(msg)
	}
}

func (serverManager *ServerManager) handleMessages() {
	for {
		msg := <-serverManager.broadcast
		serverManager.mutex.Lock()
		for client := range serverManager.usersManager.Users {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				client.Close()
				delete(serverManager.usersManager.Users, client)
			}
		}
		serverManager.mutex.Unlock()
	}
}

package main

import (
	"chat/internal/server"
)

func main() {
	serverManager := server.NewServerManager()
	serverManager.Start()
}

package main

import (
	"fgsb/internal/server"
)

func main() {
	server := server.NewServer(8080)
	server.Run()
}

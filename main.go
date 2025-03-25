package main

import (
	"embed"
	"fgsb/internal/server"
	"io/fs"
)

//go:embed web
var static embed.FS
var assets, _ = fs.Sub(static, "web/assets")

func main() {
	server.Templates = static
	server.Assets = assets
	server := server.NewServer(8080)
	server.Run()
}

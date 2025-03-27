package main

import (
	"embed"
	"fmt"
	"io/fs"

	"fyne.io/systray"

	"fgsb/internal/icon"
	"fgsb/internal/server"
)

//go:embed web
var static embed.FS
var assets, _ = fs.Sub(static, "web/assets")

func main() {
	onExit := func() {
		fmt.Println("Exit")
	}
	
	systray.Run(onReady, onExit)
}

func initMenuItem(server *server.Server){
	scoreboardItem := systray.AddMenuItem("Open Scoreboard", "Open Scoreboard")
	go func() {
		for range scoreboardItem.ClickedCh {
			server.Open("")
		}
	}()

	adminItem := systray.AddMenuItem("Admin Panel", "Open Admin Panel")
	go func() {
		for range adminItem.ClickedCh {
			server.Open("/admin/edit-scoreboard")
		}
	}()

	quitItem := systray.AddMenuItem("Quit", "Exit the app")
	go func() {
		for range quitItem.ClickedCh {
			systray.Quit()
		}
	}()
}

func onReady() {
	systray.SetTemplateIcon(icon.Icon, icon.Icon)
	systray.SetTitle("FGSB")
	systray.SetTooltip("FGSB")
	
	server.Templates = static
	server.Assets = assets
	server := server.NewServer(8080)

	go server.Run()
	
	initMenuItem(server)
	server.Open("")
}

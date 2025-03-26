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

func addQuitItem() {
	mQuit := systray.AddMenuItem("Quit", "Exit the app")
	go func() {
		for range mQuit.ClickedCh {
			systray.Quit()
		}
	}()
}

func onReady() {
	systray.SetTemplateIcon(icon.Icon, icon.Icon)
	systray.SetTitle("FGSB")
	systray.SetTooltip("FGSB")
	addQuitItem()

	server.Templates = static
	server.Assets = assets
	server := server.NewServer(8080)
	server.Run()
}

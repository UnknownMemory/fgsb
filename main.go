package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"fyne.io/systray"

	"fgsb/internal/icon"
	"fgsb/internal/server"
)

type Config struct {
	Theme string `json:"theme"`
	Port int `json:"port"`
}

//go:embed web
var static embed.FS
var assets, _ = fs.Sub(static, "web/assets")


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
	configFile, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
	}

	serv := server.Server{}
	err = json.Unmarshal(configFile, &serv)
	if err != nil {
		fmt.Println(err)
	}

	server.Templates = static
	server.Assets = assets
	go serv.Run()
			
	systray.SetTemplateIcon(icon.Icon, icon.Icon)
	systray.SetTitle("FGSB")
	systray.SetTooltip("FGSB")
	initMenuItem(&serv)
	
	serv.Open("")
}

func main() {
	onExit := func() {
		fmt.Println("Exit")
	}
	
	systray.Run(onReady, onExit)
}

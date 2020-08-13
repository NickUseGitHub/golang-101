package main

import (
	"github.com/NickUseGitHub/golang-101/configs"

	"github.com/NickUseGitHub/golang-101/app"
)

func main() {
	apiApp := &app.App{}
	apiApp.Initialize((configs.Configs{}).Initialize())
	apiApp.Run()
}

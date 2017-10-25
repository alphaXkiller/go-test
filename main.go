package main

import (
	"rr/research-tool/config"
)

func main() {
	app := config.App{}

	app.Initialize()
	app.Router.Run()
}

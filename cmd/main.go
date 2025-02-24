package main

import (
	"github.com/vaaxooo/tgaccessbot/internal/app"
	"github.com/vaaxooo/tgaccessbot/internal/bot"
)

func main() {
	application := app.NewApp()
	bot.Start(application)
}

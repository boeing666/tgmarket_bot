package main

import (
	"tgmarket/internal/app"
	"tgmarket/internal/bot"
	"tgmarket/internal/config"
	"tgmarket/internal/database"
)

func main() {
	config, err := config.Init()
	if err != nil {
		panic(err)
	}

	db, err := database.Init(config.GetDatabaseQuery())
	if err != nil {
		panic(err)
	}

	app := app.GetContainer()
	app.Init(db)

	err = bot.Run(config.APIToken)
	if err != nil {
		panic(err)
	}
}

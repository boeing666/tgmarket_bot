package main

import (
	"fmt"
	"tgmarket/internal/parser"

	"github.com/k0kubun/pp"
)

func main() {
	/*config, err := config.Init()
	if err != nil {
		panic(err)
	}

	db, err := database.Init(config.GetDatabaseQuery())
	if err != nil {
		panic(err)
	}

	app := app.GetApp()
	app.Database = db

	err = bot.Run(config.APIToken)
	if err != nil {
		panic(err)
	}*/

	yandex := parser.YandexMarket()
	info, err := yandex.GetProductInfo("https://market.yandex.ru/product--palto-coressi/621359582")
	if err != nil {
		fmt.Println(err)
	}
	pp.Println(info)
}

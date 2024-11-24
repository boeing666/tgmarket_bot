package main

import (
	"fmt"
	"tgmarket/internal/parser"
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
	product, err := parser.OZ().GetProductInfo("https://www.ozon.ru/product/organik-logos-vitamins-808638694/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*product)

	product, err = parser.OZ().GetProductInfo("https://www.ozon.ru/product/kitfort-elektricheskiy-chaynik-kt-6192-chernyy-1666181453/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*product)
}

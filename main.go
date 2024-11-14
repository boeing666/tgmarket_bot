package main

import (
	"fmt"
	"regexp"
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

	urls := []string{
		"https://www.ozon.ru/product/karikids-snegokat-dlina-99-sm-shirina-40-sm-202134869",
		"https://www.ozon.ru/product/svitshot-1694117319/?avtc=1&avte=4&avts=1732014861",
		"https://www.ozon.ru/product/mixit-skrab-dlya-tela-i-nog-antitsellyulitnyy-i-pitatelnyy-krem-batter-ot-rastyazhek-s-maslom-1077617414/?avtc=1&avte=4&avts=1732014861",
		"https://www.ozon.ru/product/1234567890",
	}

	re := regexp.MustCompile(`^((www.)|(https://www.)|(https://))*ozon.ru/product/.*?(\d{9,})`)

	for _, url := range urls {
		match := re.FindStringSubmatch(url)
		if match != nil {
			fmt.Println(match[5])
		}
	}
}

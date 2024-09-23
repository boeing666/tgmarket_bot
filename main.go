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

	app := app.GetApp()
	app.Database = db

	err = bot.Run(config.APIToken)
	if err != nil {
		panic(err)
	}
}

/* parser for megamarket
func main() {
	data := map[string]interface{}{
		"url": "https://megamarket.ru/catalog/?q=lays",
		"auth": map[string]interface{}{
			"locationId":  "66",
			"appPlatform": "WEB",
			"appVersion":  0,
			"experiments": nil,
			"os":          "UNKNOWN_OS",
		},
	}

	headers := map[string]string{
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
		"Accept":           "application/json",
		"Sec-Fetch-Site":   "same-origin",
		"Sec-Fetch-Mode":   "cors",
		"Sec-Fetch-User":   "?1",
		"Sec-Fetch-Dest":   "empty",
		"Accept-Language":  "en",
		"Authority":        "megamarket.ru",
		"Content-Type":     "application/json",
		"Origin":           "https://megamarket.ru",
		"Referer":          "https://megamarket.ru/",
		"X-Requested-With": "XMLHttpRequest",
	}

	// Преобразуем данные в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	client := cycletls.Init()

	response, err := client.Do("https://megamarket.ru/api/mobile/v1/urlService/url/parse", cycletls.Options{
		Body:      string(jsonData),
		Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,17513-27-5-65281-11-16-45-13-51-10-65037-35-43-23-18-0-41,29-23-24,0",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
		Headers:   headers,
	}, "POST")
	if err != nil {
		log.Print("Request Failed: " + err.Error())
	}
	log.Println(response)
}
*/

package parser

import "github.com/Danny-Dasilva/CycleTLS/cycletls"

type YandexParser struct {
	client cycletls.CycleTLS
}

var yandexParser *YandexParser

func YandexMarket() *YandexParser {
	if yandexParser == nil {
		yandexParser = &YandexParser{
			client: cycletls.Init(),
		}
	}
	return yandexParser
}

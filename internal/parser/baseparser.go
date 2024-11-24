package parser

import (
	"errors"
	"fmt"
)

type MarketProduct struct {
	ID      int
	Price   int
	Bonuses int
	Title   string
}

var (
	mm     = MM()
	ozon   = OZ()
	yandex = Yandex()
	wb     = WB()
)

func GetProductInfo(url string) (*MarketProduct, error) {
	parsers := []Parser{&mm, &ozon, &yandex, &wb}
	for _, parser := range parsers {
		product, err := parser.GetProductInfo(url)
		if err != nil {
			fmt.Printf("%s %v", url, err)
			continue
		}
		return product, nil
	}
	return nil, errors.New("product not found")
}

type Parser interface {
	GetProductInfo(url string) (*MarketProduct, error)
}

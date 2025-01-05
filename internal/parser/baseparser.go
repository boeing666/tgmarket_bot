package parser

import (
	"errors"
)

type MarketProduct struct {
	ID      int
	Price   int
	Bonuses int
	Title   string
}

var (
	mm = MM()
	//ozon   = OZ()
	yandex = Yandex()
	wb     = WB()
)

func GetProductInfo(url string) (*MarketProduct, error) {
	parsers := []Parser{&mm, &yandex, &wb}
	for _, parser := range parsers {
		product, err := parser.GetProductInfo(url)
		if err == nil {
			return product, nil
		}

	}
	return nil, errors.New("product not found")
}

type Parser interface {
	GetProductInfo(url string) (*MarketProduct, error)
}

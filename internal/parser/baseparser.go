package parser

type MarketProduct struct {
	Price   int
	Bonuses int
	Title   string
}

type Parser interface {
	GetProductInfo(url string, err error) *MarketProduct
}

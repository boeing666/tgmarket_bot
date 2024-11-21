package parser

type MarketProduct struct {
	ID      int
	Price   int
	Bonuses int
	Title   string
}

type Parser interface {
	GetProductInfo(url string) (*MarketProduct, error)
}

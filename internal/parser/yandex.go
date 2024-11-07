package parser

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type YandexParser struct{}

func Yandex() *YandexParser {
	y := YandexParser{}
	return &y
}

func (c YandexParser) GetProductInfo(url string, err error) (*MarketProduct, error) {
	res, err := request(url, nil, "", "GET")
	if err != nil {
		return nil, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(res.Body))
	if err != nil {
		return nil, err
	}

	product, err := findOfferInfo(doc)

	return product, err
}

func findOfferInfo(doc *html.Node) (*MarketProduct, error) {
	title := htmlquery.Find(doc, "//*[@data-auto=\"productCardTitle\"]")
	if len(title) == 0 {
		return nil, errors.New("not found productCardTitle")
	}

	offerHTML := htmlquery.Find(doc, "//div[@data-zone-name='cpa-offer']")
	if len(offerHTML) == 0 {
		return nil, errors.New("not found cpa-offer")
	}

	var offerData OfferYM
	offer := htmlquery.SelectAttr(offerHTML[0], "data-zone-data")
	err := json.Unmarshal([]byte(offer), &offerData)
	if err != nil {
		return nil, err
	}

	var product MarketProduct
	product.Title = htmlquery.InnerText(title[0])
	product.Price = offerData.PriceDetails.Price.Value

	return &product, nil
}

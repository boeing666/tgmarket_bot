package parser

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/Noooste/azuretls-client"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type YandexParser struct {
	reId regexp.Regexp
}

func Yandex() YandexParser {
	return YandexParser{
		reId: *regexp.MustCompile(`^https://market\.yandex\.ru/product--.*?\/([0-9]+)(\?.*)?$`),
	}
}

func (y YandexParser) GetProductInfo(url string) (*MarketProduct, error) {
	f := y.reId.FindStringSubmatch(url)
	if len(f) < 2 {
		return nil, errors.New("can't find item id")
	}

	session := azuretls.NewSession()
	defer session.Close()

	res, err := session.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(string(res.Body)))
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
	product.ID = offerData.ProductID

	return &product, nil
}

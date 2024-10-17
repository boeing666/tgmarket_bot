package parser

import (
	"encoding/json"
	"strings"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type YandexParser struct {
	client cycletls.CycleTLS
}
type YMProduct struct {
	Price int
	Title string
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

func (c YandexParser) GetProductInfo(url string) (*YMProduct, error) {
	res, err := c.request(url, "")
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

func (c YandexParser) request(url string, body string) (cycletls.Response, error) {
	return c.client.Do(url, cycletls.Options{
		Body:      body,
		Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,17513-27-5-65281-11-16-45-13-51-10-65037-35-43-23-18-0-41,29-23-24,0",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
	}, "GET")
}

func findOfferInfo(doc *html.Node) (*YMProduct, error) {
	title := htmlquery.Find(doc, "//*[@data-auto=\"productCardTitle\"]")
	if len(title) == 0 {
		return nil, nil
	}

	offerHTML := htmlquery.Find(doc, "//div[@data-zone-name='cpa-offer']")
	if len(offerHTML) == 0 {
		return nil, nil
	}

	var offerData OfferYM
	offer := htmlquery.SelectAttr(offerHTML[0], "data-zone-data")
	err := json.Unmarshal([]byte(offer), &offerData)
	if err != nil {
		return nil, nil
	}

	var info YMProduct
	info.Title = htmlquery.InnerText(title[0])
	info.Price = offerData.PriceDetails.Price.Value

	return &info, nil
}

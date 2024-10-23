package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

type Ozon struct {
	urlTemplate string
	regexID     regexp.Regexp
}

func OZ() *Ozon {
	p := Ozon{
		urlTemplate: "https://www.ozon.ru/api/composer-api.bx/page/json/v2?url=%s",
		regexID:     *regexp.MustCompile(`^((www.)|(https://www.)|(https://))*ozon.ru/product/.*?(\d{9,})`),
	}
	return &p
}

func (m Ozon) GetProductInfo(url string, err error) (*MarketProduct, error) {
	apiurl := fmt.Sprintf(m.urlTemplate, url)

	res, err := request(apiurl, nil, "", "GET")
	if err != nil {
		return nil, err
	}

	if res.Status != 200 {
		return nil, fmt.Errorf("error. status code: %v", res.Status)
	}

	var wbres WBResponse
	err = json.Unmarshal([]byte(res.Body), &wbres)
	if err != nil {
		return nil, err
	}

	if len(wbres.Data.Products) == 0 {
		return nil, errors.New("field \"product\" is empty.")
	}

	wbProduct := wbres.Data.Products[0]

	var product MarketProduct
	product.Title = wbProduct.Name
	product.Price = wbProduct.Sizes[0].Price.Total / 100.0

	return &product, err
}

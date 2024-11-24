package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

type Wildberries struct {
	urlTemplate string
	reId        regexp.Regexp
}

func Wb() *Wildberries {
	p := Wildberries{
		urlTemplate: "https://card.wb.ru/cards/v2/detail?appType=1&curr=rub&dest=-5818883&ab_testing=false&nm=%s",
		reId:        *regexp.MustCompile(`^((www.)|(https://www.)|(https://))*wildberries.ru/catalog/(\d+)\S*\z`),
	}
	return &p
}

func (w Wildberries) GetProductInfo(url string) (*MarketProduct, error) {
	f := w.reId.FindStringSubmatch(url)
	if len(f) < 2 {
		return nil, errors.New("can't find item id")
	}

	itemId := f[len(f)-1]
	apiurl := fmt.Sprintf(w.urlTemplate, itemId)

	res, err := tlsRequest(apiurl, nil, "", "GET")
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
	product.ID = wbProduct.ID

	return &product, err
}

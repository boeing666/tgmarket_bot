package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type Ozon struct {
	urlTemplate string
	reId        regexp.Regexp
}

func OZ() Ozon {
	return Ozon{
		urlTemplate: "https://www.ozon.ru/api/composer-api.bx/page/json/v2?url=%s",
		reId:        *regexp.MustCompile(`^((www.)|(https://www.)|(https://))*ozon.ru/product/.*?(\d{9,})`),
	}
}

func (o Ozon) GetProductInfo(url string) (*MarketProduct, error) {
	f := o.reId.FindStringSubmatch(url)
	if len(f) < 2 {
		return nil, errors.New("can't find item id")
	}

	apiurl := fmt.Sprintf(o.urlTemplate, url)

	resp, err := httpClient().Get(apiurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resp, err = httpClient().Get(apiurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var jsonRes OzonInfo
	err = json.Unmarshal(body, &jsonRes)
	if err != nil {
		return nil, err
	}

	if len(jsonRes.Seo.Script) == 0 {
		return nil, errors.New("empty innerHTML")
	}

	var jsonProduct OzonProductInfo
	err = json.Unmarshal([]byte(jsonRes.Seo.Script[0].InnerHTML), &jsonProduct)
	if err != nil {
		return nil, err
	}

	var product MarketProduct
	product.Title = jsonProduct.Name
	product.Price, _ = strconv.Atoi(jsonProduct.Offers.Price)
	product.Bonuses = 0
	product.ID, _ = strconv.Atoi(jsonProduct.Sku)

	return &product, err
}

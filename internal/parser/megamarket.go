package parser

import (
	"encoding/json"
	"regexp"
)

type MegaMarket struct{}

func MM() *MegaMarket {
	m := MegaMarket{}
	return &m
}

func getAuth() map[string]any {
	return map[string]any{
		"locationId":  "50",
		"appPlatform": "WEB",
		"appVersion":  1710405202,
		"experiments": nil,
		"os":          "UNKNOWN_OS",
	}
}

func getHeaders() map[string]string {
	return map[string]string{
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
		"Accept":           "application/json",
		"Sec-Fetch-Site":   "same-origin",
		"Sec-Fetch-Mode":   "cors",
		"Sec-Fetch-User":   "?1",
		"Sec-Fetch-Dest":   "empty",
		"Accept-Language":  "en",
		"Authority":        "megamarket.ru",
		"Content-Type":     "application/json",
		"Origin":           "https://megamarket.ru",
		"Referer":          "https://megamarket.ru/",
		"X-Requested-With": "XMLHttpRequest",
	}
}

func getOffersForProduct(goodsId string) string {
	data := map[string]interface{}{
		"addressId":    nil,
		"collectionId": nil,
		"goodsId":      goodsId,
		"listingParams": map[string]interface{}{
			"priorDueDate":    "UNKNOWN_OFFER_DUE_DATE",
			"selectedFilters": nil,
		},
		"merchantId":     "0",
		"requestVersion": 11,
		"shopInfo":       nil,
		"auth":           getAuth(),
	}
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func generateJsonForProduct(goodsId string) string {
	data := map[string]interface{}{
		"goodsId":    goodsId,
		"merchantId": "0",
		"auth":       getAuth(),
	}
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func (c MegaMarket) getOffers(goodsId string) (*ProductOffers, error) {
	res, err :=
		request("https://megamarket.ru/api/mobile/v1/catalogService/productOffers/get",
			getHeaders(),
			getOffersForProduct(goodsId),
			"POST")
	if err != nil {
		return nil, err
	}

	var data ProductOffers
	err = json.Unmarshal([]byte(res.Body), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c MegaMarket) GetProductInfo(goodsId string) (*MarketProduct, error) {
	res, err :=
		request("https://megamarket.ru/api/mobile/v1/catalogService/productCardMainInfo/get",
			getHeaders(),
			generateJsonForProduct(goodsId),
			"POST")
	if err != nil {
		return nil, err
	}

	var data ProductInfo
	err = json.Unmarshal([]byte(res.Body), &data)
	if err != nil {
		return nil, err
	}

	offers, err := c.getOffers(goodsId)
	if err != nil {
		return nil, err
	}

	price, bonus := findLowestPriceAndHighBonusesMM(offers)

	var product MarketProduct
	product.Title = data.Goods.Title
	product.Price = price
	product.Bonuses = bonus

	return &product, nil
}

func findLowestPriceAndHighBonusesMM(offers *ProductOffers) (int, int) {
	if !offers.IsAvailable {
		return 0, 0
	}

	if len(offers.Offers) == 0 {
		return 0, 0
	}

	lowestPrice := offers.Offers[0].FinalPrice
	highestBonuses := offers.Offers[0].BonusAmountFinalPrice

	for _, offer := range offers.Offers {
		if lowestPrice >= offer.FinalPrice {
			lowestPrice = offer.FinalPrice
		}

		if highestBonuses < offer.BonusAmountFinalPrice {
			highestBonuses = offer.BonusAmountFinalPrice
		}
	}

	return lowestPrice, highestBonuses
}

func GetProductIDFromUrl(url string) (string, bool) {
	regex := regexp.MustCompile(`(\d{12})(?:_(\d{5}))?`)
	res := regex.FindStringSubmatch(url)
	if len(res) == 0 {
		return "", false
	}

	return res[0], true
}

package parser

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"

	"github.com/Noooste/azuretls-client"
)

type MegaMarket struct {
	reId regexp.Regexp
}

func MM() MegaMarket {
	return MegaMarket{
		reId: *regexp.MustCompile(`^https:\/\/megamarket\.ru\/catalog\/details\/.*?-(\d+)(?:\/|$)`),
	}
}

func getAuth() map[string]any {
	return map[string]any{
		"locationId":  "50",
		"appPlatform": "WEB",
		"appVersion":  1710405202,
		"experiments": map[string]interface{}{},
		"os":          "UNKNOWN_OS",
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

func (m MegaMarket) getOffers(session *azuretls.Session, goodsId string) (*ProductOffers, error) {
	res, err := session.Post("https://megamarket.ru/api/mobile/v1/catalogService/productOffers/get",
		getOffersForProduct(goodsId))

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

func (m MegaMarket) GetProductInfo(url string) (*MarketProduct, error) {
	f := m.reId.FindStringSubmatch(url)
	if len(f) < 2 {
		return nil, errors.New("can't find item id")
	}

	goodsId := f[len(f)-1]
	goodsIdInt, err := strconv.Atoi(goodsId)
	if err != nil {
		return nil, err
	}

	session := azuretls.NewSession()
	defer session.Close()

	res, err := session.Post("https://megamarket.ru/api/mobile/v1/catalogService/productCardMainInfo/get", generateJsonForProduct(goodsId))
	if err != nil {
		return nil, err
	}

	var data ProductInfo
	err = json.Unmarshal([]byte(res.Body), &data)
	if err != nil {
		return nil, err
	}

	if !data.Success {
		return nil, errors.New("api success false")
	}

	offers, err := m.getOffers(session, goodsId)
	if err != nil {
		return nil, err
	}

	price, bonus := findLowestPriceAndHighBonusesMM(offers)

	var product MarketProduct
	product.Title = data.Goods.Title
	product.Price = price
	product.Bonuses = bonus
	product.ID = goodsIdInt

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

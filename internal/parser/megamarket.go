package parser

import (
	"encoding/json"
	"regexp"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

type MegaParser struct {
	client cycletls.CycleTLS
}

var megamarket *MegaParser

func MegaMarket() *MegaParser {
	if megamarket == nil {
		megamarket = &MegaParser{
			client: cycletls.Init(),
		}
	}
	return megamarket
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

func getProductInfo(goodsId string) string {
	data := map[string]interface{}{
		"goodsId":    goodsId,
		"merchantId": "0",
		"auth":       getAuth(),
	}
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func (c MegaParser) GetOffers(goodsId string) (*ProductOffers, error) {
	res, err := c.request("https://megamarket.ru/api/mobile/v1/catalogService/productOffers/get", getOffersForProduct(goodsId))
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

func (c MegaParser) GetProductInfo(goodsId string) (*ProductInfo, error) {
	res, err := c.request("https://megamarket.ru/api/mobile/v1/catalogService/productCardMainInfo/get", getProductInfo(goodsId))
	if err != nil {
		return nil, err
	}

	var data ProductInfo
	err = json.Unmarshal([]byte(res.Body), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c MegaParser) request(url string, body string) (cycletls.Response, error) {
	return c.client.Do(url, cycletls.Options{
		Headers:   getHeaders(),
		Body:      body,
		Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,17513-27-5-65281-11-16-45-13-51-10-65037-35-43-23-18-0-41,29-23-24,0",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
	}, "POST")
}

func GetProductIDFromUrl(url string) (string, bool) {
	regex := regexp.MustCompile(`(\d{12})(?:_(\d{5}))?`)
	res := regex.FindStringSubmatch(url)
	if len(res) == 0 {
		return "", false
	}
	return res[0], true
}

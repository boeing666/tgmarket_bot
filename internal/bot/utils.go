package bot

import (
	"net/url"
	"strings"
	"tgmarket/internal/protobufs"
)

func detectMarketplace(link string) protobufs.Shops {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return protobufs.Shops_UnknownShop
	}

	switch {
	case strings.Contains(parsedURL.Host, "market.yandex.ru"):
		return protobufs.Shops_YandexMarket
	case strings.Contains(parsedURL.Host, "ozon.ru"):
		return protobufs.Shops_Ozon
	case strings.Contains(parsedURL.Host, "wildberries.ru"):
		return protobufs.Shops_Wildberries
	case strings.Contains(parsedURL.Host, "megamarket.ru"):
		return protobufs.Shops_SberMegaMarket
	}

	return protobufs.Shops_UnknownShop
}

func getShopName(shop protobufs.Shops) string {
	switch shop {
	case protobufs.Shops_YandexMarket:
		return "Яндекс Маркет"
	case protobufs.Shops_Ozon:
		return "Озон"
	case protobufs.Shops_Wildberries:
		return "Wildberries"
	case protobufs.Shops_SberMegaMarket:
		return "МегаМаркет"
	default:
		return "Неизвестный"
	}
}

package bot

import (
	"fmt"
	"math"
	"net/url"
	"sort"
	"strings"
	"tgmarket/internal/cache"
	"tgmarket/internal/models"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	"google.golang.org/protobuf/proto"
)

type menuInfo struct {
	page     int
	maxPages int
}

func detectMarketplace(link string) protobufs.Shops {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return protobufs.Shops_UnknownShop
	}

	switch {
	/*case strings.Contains(parsedURL.Host, "market.yandex.ru"):
		return protobufs.Shops_YandexMarket
	case strings.Contains(parsedURL.Host, "ozon.ru"):
		return protobufs.Shops_Ozon
	case strings.Contains(parsedURL.Host, "wildberries.ru"):
		return protobufs.Shops_Wildberries*/
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

func filterUserProducts(user *cache.User, name string) map[int64]*models.Product {
	lowername := strings.ToLower(name)
	filteredProducts := make(map[int64]*models.Product)
	for _, product := range user.Products {
		if strings.Contains(strings.ToLower(product.Name), lowername) {
			filteredProducts[product.ID] = product
		}
	}
	return filteredProducts
}

func buildMenuHeader(user *cache.User, menu menuInfo) [][]telego.InlineKeyboardButton {
	var rows [][]telego.InlineKeyboardButton

	if len(user.FilterName) == 0 {
		rows = append(rows, protobufs.CreateRowButton("🔎 Поиск товара по имени", protobufs.ButtonID_SearchByName, nil))
	} else {
		text := fmt.Sprintf("🔎 Убрать поиск по имени (%s)", user.FilterName)
		rows = append(rows, protobufs.CreateRowButton(text, protobufs.ButtonID_RemoveFilterByName, nil))
	}

	if menu.maxPages == 0 {
		rows = append(rows, protobufs.CreateRowButton("🥺 Список товаров пуст, добавьте что-нибудь", protobufs.ButtonID_Nothing, nil))
	} else {
		text := fmt.Sprintf("📄 Страница %d/%d", menu.page+1, menu.maxPages)
		rows = append(rows, protobufs.CreateRowButton(text, protobufs.ButtonID_Nothing, nil))
	}

	return rows
}

func buildPage[V any](curpage int, data map[int64]V) ([]int64, menuInfo) {
	maxPages := 0
	pageSize := 5

	maxPages = int(math.Ceil(float64(len(data)) / float64(pageSize)))

	if curpage < 0 {
		curpage = 0
	} else if curpage >= maxPages {
		curpage = maxPages - 1
	}

	startIndex := curpage * pageSize
	endIndex := startIndex + pageSize

	keys := make([]int64, 0, len(data))

	for k := range data {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	values := make([]int64, 0, pageSize)

	i := 0
	for _, id := range keys {
		if i >= startIndex && i < endIndex {
			values = append(values, id)
		}
		i++
	}

	return values, menuInfo{page: curpage, maxPages: maxPages}
}

func buildNavigation(menu menuInfo, messageID protobufs.ButtonID) []telego.InlineKeyboardButton {
	var buttons []telego.InlineKeyboardButton

	if menu.page > 0 {
		buttons = append(buttons,
			protobufs.CreateButton("⏪", messageID, &protobufs.ChangePage{Newpage: proto.Int64(0)}),
			protobufs.CreateButton("⬅️", messageID, &protobufs.ChangePage{Newpage: proto.Int64(int64(menu.page - 1))}),
		)
	}

	if menu.page < menu.maxPages-1 {
		buttons = append(buttons,
			protobufs.CreateButton("➡️", messageID, &protobufs.ChangePage{Newpage: proto.Int64(int64(menu.page + 1))}),
			protobufs.CreateButton("⏩", messageID, &protobufs.ChangePage{Newpage: proto.Int64(int64(menu.maxPages))}),
		)
	}

	return buttons
}

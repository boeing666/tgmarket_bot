package bot

import (
	"fmt"
	"tgmarket/internal/cache"
	"tgmarket/internal/models"
	"tgmarket/internal/parser"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func sendMinPriceMessage(user *cache.User, bot *telego.Bot, product *models.Product) {
	text, entities := tu.MessageEntities(
		tu.Entity("🎉 Ура! Цена на товар стала ниже минимальной! 🎉\n"),
		tu.Entity("📦 Товар: "), tu.Entity(fmt.Sprintf("%s\n", product.Name)).TextLink(product.URL),
		tu.Entity(fmt.Sprintf("📉 Текущая цена: %d\n", product.Price)),
		tu.Entity(fmt.Sprintf("💰 Минимальная цена: %d\n", product.MinPrice)),
		tu.Entity("Не упустите возможность купить! 💸\n"),
	)

	bot.SendMessage(tu.Message(
		tu.ID(user.TelegramID),
		text,
	).WithEntities(entities...))
}

func sendMinBonusesMessage(user *cache.User, bot *telego.Bot, product *models.Product, bonus int) {
	text, entities := tu.MessageEntities(
		tu.Entity("🎉 Бонусов за товар стало больше! 🎉\n"),
		tu.Entity("📦 Товар: "), tu.Entity(fmt.Sprintf("%s\n", product.Name)).TextLink(product.URL),
		tu.Entity(fmt.Sprintf("🏆 Бонусов: %d\n", bonus)),
		tu.Entity(fmt.Sprintf("📈 Минимальное количество бонусов: %d\n", product.MinBonuses)),
		tu.Entity("Не упустите шанс получить больше выгоды! 🌟\n"),
	)

	bot.SendMessage(tu.Message(
		tu.ID(user.TelegramID),
		text,
	).WithEntities(entities...))
}

func findLowestPriceAndHighBonuses(offers *parser.ProductOffers) (int, int) {
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

func marketParser(bot *telego.Bot) {
	mm := parser.MegaMarket()
	for {
		for _, user := range userscache.users {
			for _, product := range user.Products {
				productOffers, err := mm.GetOffers(product.ProductID)
				if err != nil || !productOffers.IsAvailable {
					continue
				}
				price, bonus := findLowestPriceAndHighBonuses(productOffers)

				if price != product.Price {
					product.Price = price
					if product.MinPrice >= price {
						sendMinPriceMessage(user, bot, product)
					}
				}

				if bonus != product.Bonus {
					product.Bonus = bonus
					if product.MinBonuses <= bonus {
						sendMinBonusesMessage(user, bot, product, bonus)
					}
				}

				user.UpdateProduct(product)
				time.Sleep(5 * time.Second)
			}
			time.Sleep(time.Minute)
		}
		time.Sleep(time.Minute)
	}
}

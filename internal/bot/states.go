package bot

import (
	"fmt"
	"strconv"
	"tgmarket/internal/cache"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func handleEnterProductURL(user *cache.User, bot *telego.Bot, update *telego.Update) error {
	bot.DeleteMessage(tu.Delete(tu.ID(update.Message.Chat.ID), update.Message.MessageID))

	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(protobufs.ButtonCancel()),
	)

	market := detectMarketplace(update.Message.Text)
	if market == protobufs.Shops_UnknownShop {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:      tu.ID(update.Message.Chat.ID),
			Text:        errorInProductURLText(),
			MessageID:   user.ActiveMsgID,
			ReplyMarkup: keyboard,
		})
		return nil
	}

	product, err := user.AddProduct(int(market), update.Message.Text)
	if err != nil {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:      tu.ID(update.Message.Chat.ID),
			Text:        errorAddProductToDBText(),
			MessageID:   user.ActiveMsgID,
			ReplyMarkup: keyboard,
		})
		return err
	}

	user.State = protobufs.UserState_None
	return showProductInfo(product.ID, bot, user)
}

func handleEnterProductName(user *cache.User, bot *telego.Bot, update *telego.Update) error {
	if err := bot.DeleteMessage(tu.Delete(tu.ID(update.Message.Chat.ID), update.Message.MessageID)); err != nil {
		return err
	}

	product, found := user.Products[user.ActiveProductID]
	if !found {
		return fmt.Errorf("EnterProductName id %d user %d", product.ID, user.TelegramID)
	}

	product.Name = update.Message.Text
	if err := user.UpdateProduct(product); err != nil {
		return err
	}

	return showProductInfo(product.ID, bot, user)
}

func handleEnterMinPrice(user *cache.User, bot *telego.Bot, update *telego.Update) error {
	if err := bot.DeleteMessage(tu.Delete(tu.ID(update.Message.Chat.ID), update.Message.MessageID)); err != nil {
		return err
	}

	product, found := user.Products[user.ActiveProductID]
	if !found {
		return fmt.Errorf("EnterMinPrice id %d user %d", product.ID, user.TelegramID)
	}
	value, _ := strconv.Atoi(update.Message.Text)
	product.MinPrice = uint(value)

	if err := user.UpdateProduct(product); err != nil {
		return err
	}

	return showProductInfo(product.ID, bot, user)
}

func handleEnterMinBonuses(user *cache.User, bot *telego.Bot, update *telego.Update) error {
	if err := bot.DeleteMessage(tu.Delete(tu.ID(update.Message.Chat.ID), update.Message.MessageID)); err != nil {
		return err
	}

	product, found := user.Products[user.ActiveProductID]
	if !found {
		return fmt.Errorf("MinBonuses id %d user %d", product.ID, user.TelegramID)
	}

	value, _ := strconv.Atoi(update.Message.Text)
	product.MinBonuses = uint(value)

	if err := user.UpdateProduct(product); err != nil {
		return err
	}

	return showProductInfo(product.ID, bot, user)
}

func handleUserStates(bot *telego.Bot, update telego.Update) {
	ctx := update.Context()
	user := ctx.Value("user").(*cache.User)

	if user.State == protobufs.UserState_None {
		return
	}

	var err error
	switch user.State {
	case protobufs.UserState_EnterProductURL:
		err = handleEnterProductURL(user, bot, &update)
	case protobufs.UserState_EnterProductName:
		err = handleEnterProductName(user, bot, &update)
	case protobufs.UserState_EnterMinPrice:
		err = handleEnterMinPrice(user, bot, &update)
	case protobufs.UserState_EnterMinBonuses:
		err = handleEnterMinBonuses(user, bot, &update)
	}

	if err != nil {
		fmt.Println(err)
	}
}

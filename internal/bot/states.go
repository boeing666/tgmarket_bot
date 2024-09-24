package bot

import (
	"context"
	"fmt"
	"strconv"
	"tgmarket/internal/cache"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type stateHandler func(data messageContext) error

var stateCallbacks map[protobufs.UserState]stateHandler

func handleEnterProductURL(data messageContext) error {
	user := data.GetUser()
	bot := data.GetBot()
	chatid := tu.ID(user.TelegramID)

	if err := bot.DeleteMessage(tu.Delete(chatid, data.GetMessageID())); err != nil {
		return err
	}

	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(protobufs.ButtonCancel()),
	)

	market := detectMarketplace(data.GetMessageText())
	if market == protobufs.Shops_UnknownShop {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:      chatid,
			Text:        errorInProductURLText(),
			MessageID:   user.ActiveMsgID,
			ReplyMarkup: keyboard,
		})
		return nil
	}

	product, err := user.AddProduct(int(market), data.GetMessageText())
	if err != nil {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:      chatid,
			Text:        errorAddProductToDBText(),
			MessageID:   user.ActiveMsgID,
			ReplyMarkup: keyboard,
		})
		return err
	}

	user.State = protobufs.UserState_None
	return showProductInfo(product.ID, bot, user)
}

func handleEnterProductName(data messageContext) error {
	user := data.GetUser()
	bot := data.GetBot()

	if err := bot.DeleteMessage(tu.Delete(tu.ID(user.TelegramID), data.GetMessageID())); err != nil {
		return err
	}

	product, found := user.Products[user.ActiveProductID]
	if !found {
		return fmt.Errorf("EnterProductName id %d user %d", product.ID, user.TelegramID)
	}

	product.Name = data.GetMessageText()
	if err := user.UpdateProduct(product); err != nil {
		return err
	}

	return showProductInfo(product.ID, bot, user)
}

func handleEnterMinPrice(data messageContext) error {
	user := data.GetUser()
	bot := data.GetBot()

	if err := bot.DeleteMessage(tu.Delete(tu.ID(user.TelegramID), data.GetMessageID())); err != nil {
		return err
	}

	product, found := user.Products[user.ActiveProductID]
	if !found {
		return fmt.Errorf("EnterMinPrice id %d user %d", product.ID, user.TelegramID)
	}

	value, _ := strconv.Atoi(data.GetMessageText())
	product.MinPrice = uint(value)

	if err := user.UpdateProduct(product); err != nil {
		return err
	}

	return showProductInfo(product.ID, bot, user)
}

func handleEnterMinBonuses(data messageContext) error {
	user := data.GetUser()
	bot := data.GetBot()

	if err := bot.DeleteMessage(tu.Delete(tu.ID(user.TelegramID), data.GetMessageID())); err != nil {
		return err
	}

	product, found := user.Products[user.ActiveProductID]
	if !found {
		return fmt.Errorf("MinBonuses id %d user %d", product.ID, user.TelegramID)
	}

	value, _ := strconv.Atoi(data.GetMessageText())
	product.MinBonuses = uint(value)

	if err := user.UpdateProduct(product); err != nil {
		return err
	}

	return showProductInfo(product.ID, bot, user)
}

func handleUserStates(ctx context.Context, bot *telego.Bot, message telego.Message) {
	user := ctx.Value("user").(*cache.User)
	if user.State == protobufs.UserState_None {
		return
	}

	if callback, ok := stateCallbacks[user.State]; ok {
		data := messageData{
			baseContext: baseContext{
				bot:  bot,
				user: user,
			},
			messageid:   message.MessageID,
			messagetext: message.Text,
		}
		if err := callback(&data); err != nil {
			fmt.Println(err)
		}
	}
}

func registerStates() {
	stateCallbacks = make(map[protobufs.UserState]stateHandler)

	stateCallbacks[protobufs.UserState_EnterProductURL] = handleEnterProductURL
	stateCallbacks[protobufs.UserState_EnterProductName] = handleEnterProductName
	stateCallbacks[protobufs.UserState_EnterMinPrice] = handleEnterMinPrice
	stateCallbacks[protobufs.UserState_EnterMinBonuses] = handleEnterMinBonuses
}

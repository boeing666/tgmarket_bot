package bot

import (
	"context"
	"fmt"
	"strconv"
	"tgmarket/internal/cache"
	"tgmarket/internal/parser"
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

	button := protobufs.ButtonCancel(protobufs.ButtonID_MainMenu, nil)
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(button),
	)

	market := detectMarketplace(data.GetMessageText())
	if market == protobufs.Shops_UnknownShop {
		_, err := bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:      chatid,
			Text:        errorInProductURLText(),
			MessageID:   user.ActiveMsgID,
			ReplyMarkup: keyboard,
		})
		return err
	}

	productInfo, err := parser.GetProductInfo(data.GetMessageText())
	if err != nil {
		_, err := bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:      chatid,
			Text:        errorInProductURLText(),
			MessageID:   user.ActiveMsgID,
			ReplyMarkup: keyboard,
		})
		return err
	}

	productID := strconv.Itoa(productInfo.ID)
	product := user.FindProductByProductID(productID)
	if product != nil {
		_, err := bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:      chatid,
			Text:        errorProductAlreadyExist(),
			MessageID:   user.ActiveMsgID,
			ReplyMarkup: keyboard,
		})
		return err
	}

	product, err = user.AddProduct(int(market), data.GetMessageText(), productInfo)
	if err != nil {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:      chatid,
			Text:        errorAddProductToDBText(),
			MessageID:   user.ActiveMsgID,
			ReplyMarkup: keyboard,
		})
		return err
	}

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
	product.MinPrice = value

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
	product.MinBonuses = value

	if err := user.UpdateProduct(product); err != nil {
		return err
	}

	return showProductInfo(product.ID, bot, user)
}

func handleEnterProductNameForSearch(data messageContext) error {
	user := data.GetUser()
	bot := data.GetBot()

	if err := bot.DeleteMessage(tu.Delete(tu.ID(user.TelegramID), data.GetMessageID())); err != nil {
		return err
	}

	user.FilterName = data.GetMessageText()
	user.FiltredProducts = filterUserProducts(user, data.GetMessageText())

	return callbackListOfProducts(data)
}

func handleUserStates(ctx context.Context, bot *telego.Bot, message telego.Message) {
	user := ctx.Value("user").(*cache.User)
	if user.State == protobufs.UserState_None {
		bot.SendMessage(tu.Message(
			tu.ID(user.TelegramID),
			"Используйте /start, для начала работы с ботом",
		))
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
	stateCallbacks[protobufs.UserState_EnterMinBonuses] = handleEnterMinBonuses
	stateCallbacks[protobufs.UserState_EnterPartialProductName] = handleEnterProductNameForSearch
}

package bot

import (
	"tgmarket/internal/protobufs"
	"tgmarket/internal/user"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type callbackContext struct {
	Bot    *telego.Bot
	Update *telego.Update
	Data   []byte
}

type callbackHandler func(ctx callbackContext) error

func (c callbackContext) EditMessageText(text string) (*telego.Message, error) {
	return c.Bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:    tu.ID(c.Update.CallbackQuery.From.ID),
		Text:      text,
		MessageID: c.Update.CallbackQuery.Message.GetMessageID(),
	})
}

func (c callbackContext) EditMessageTextWithKeyboard(text string, keyboard *telego.InlineKeyboardMarkup) (*telego.Message, error) {
	return c.Bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(c.Update.CallbackQuery.From.ID),
		Text:        text,
		MessageID:   c.Update.CallbackQuery.Message.GetMessageID(),
		ReplyMarkup: keyboard,
	})
}

func (c callbackContext) GetUser() *user.Cache {
	ctx := c.Update.Context()
	return ctx.Value("user").(*user.Cache)
}

var buttonCallbacks map[protobufs.ButtonID]callbackHandler

func registerButtons() {
	buttonCallbacks = make(map[protobufs.ButtonID]callbackHandler)

	buttonCallbacks[protobufs.ButtonID_AddProduct] = callbackAddProduct
	buttonCallbacks[protobufs.ButtonID_ListOfProducts] = callbackListOfProducts
	buttonCallbacks[protobufs.ButtonID_SetMinPrice] = callbackSetMinPrice
	buttonCallbacks[protobufs.ButtonID_SetMinBonuses] = callbackSetMinBonuses
	buttonCallbacks[protobufs.ButtonID_ChangeProductName] = callbackChangeProductName
	buttonCallbacks[protobufs.ButtonID_DeleteProduct] = callbackDeleteProduct
	buttonCallbacks[protobufs.ButtonID_MainMenu] = callbackMainMenu
	buttonCallbacks[protobufs.ButtonID_ChangeMenu] = callbackChangeMenu
}

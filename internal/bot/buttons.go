package bot

import (
	"tgmarket/internal/cache"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/protobuf/proto"
)

type callbackContext struct {
	Bot    *telego.Bot
	Update *telego.Update
	Data   []byte
}

type callbackHandler func(ctx callbackContext) error

func (c callbackContext) EditLastMessageText(text string) (*telego.Message, error) {
	return c.Bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:    tu.ID(c.Update.CallbackQuery.From.ID),
		Text:      text,
		MessageID: c.Update.CallbackQuery.Message.GetMessageID(),
	})
}

func (c callbackContext) EditLastMessage(text string, keyboard *telego.InlineKeyboardMarkup) (*telego.Message, error) {
	return c.Bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(c.Update.CallbackQuery.From.ID),
		Text:        text,
		MessageID:   c.Update.CallbackQuery.Message.GetMessageID(),
		ReplyMarkup: keyboard,
	})
}

func (c callbackContext) GetUser() *cache.User {
	return getUser(c.Update)
}

func (c callbackContext) Unmarshal(msg proto.Message) error {
	return proto.Unmarshal(c.Data, msg)
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
	buttonCallbacks[protobufs.ButtonID_ChangeMenu] = callbackChangeMenu
	buttonCallbacks[protobufs.ButtonID_MainMenu] = callbackMainMenu
}

package bot

import (
	"fmt"
	"tgmarket/internal/cache"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func callbackAddProduct(ctx callbackContext) error {
	ctx.EditLastMessage(addNewProductText(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(protobufs.ButtonCancel()),
		),
	)
	ctx.GetUser().State = protobufs.UserState_EnterProductURL
	return nil
}

func callbackProductInfo(ctx callbackContext) error {
	var msg protobufs.ProdcutData
	err := ctx.Unmarshal(&msg)
	if err != nil {
		return err
	}

	return showProductInfo(&msg, ctx.Bot, ctx.GetUser())
}

func callbackListOfProducts(ctx callbackContext) error {
	return nil
}

func callbackSetMinPrice(ctx callbackContext) error {
	return nil
}

func callbackSetMinBonuses(ctx callbackContext) error {
	return nil
}

func callbackChangeProductName(ctx callbackContext) error {
	return nil
}

func callbackDeleteProduct(ctx callbackContext) error {
	return nil
}

func callbackMainMenu(ctx callbackContext) error {
	ctx.GetUser().State = protobufs.UserState_None
	_, err := ctx.EditLastMessage(welcomeText(), protobufs.BuildMainMenu())
	return err
}

func callbackChangeMenu(ctx callbackContext) error {
	var msg protobufs.ButtonData
	err := ctx.Unmarshal(&msg)
	if err != nil {
		return err
	}
	ctx.Data = msg.Data
	return buttonCallbacks[msg.Id](ctx)
}

func showProductInfo(product *protobufs.ProdcutData, bot *telego.Bot, user *cache.User) error {
	//createStr := peer.CreatedAt.Format("2006-01-02 15:04:05")

	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(protobufs.ButtonSetMinimalPrice()),
		tu.InlineKeyboardRow(protobufs.ButtonSetMinimalBonuses()),
		tu.InlineKeyboardRow(protobufs.ButtonSetProductName()),
		tu.InlineKeyboardRow(protobufs.ButtonDeleteProduct()),
		tu.InlineKeyboardRow(protobufs.ButtonBack(protobufs.ButtonID_MainMenu, nil)),
		tu.InlineKeyboardRow(protobufs.ButtonMainMenu()),
	)

	messageText := fmt.Sprintf(`
📦 Товар: Хлеб
🛒 Магазин: Пятерочка
💰 Цена: 1337 руб
💵 Мин. цена: 0
❇️ Мин. кол-во бонусов: 0
🗓️ Добавлен: 20.09.2024 14:00`)

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(user.TelegramID),
		Text:        messageText,
		MessageID:   user.LastMsgID,
		ReplyMarkup: keyboard,
	})
	return nil
}

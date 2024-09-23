package bot

import (
	"fmt"
	"tgmarket/internal/cache"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func handleEnterProductURL(user *cache.User, bot *telego.Bot, update *telego.Update) error {
	bot.DeleteMessage(tu.Delete(tu.ID(update.Message.Chat.ID), update.Message.MessageID))

	market := detectMarketplace(update.Message.Text)
	if market == protobufs.Shops_UnknownShop {
		keyboard := tu.InlineKeyboard(
			tu.InlineKeyboardRow(protobufs.ButtonCancel()),
		)

		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:      tu.ID(update.Message.Chat.ID),
			Text:        errorInProductURLText(),
			MessageID:   user.LastMsgID,
			ReplyMarkup: keyboard,
		})
		return nil
	}

	return nil
}

func handleEnterProductName(user *cache.User, bot *telego.Bot, update *telego.Update) error {
	return nil
}

func handleEnterProductPool(user *cache.User, bot *telego.Bot, update *telego.Update) error {
	return nil
}

func handleEnterMinPrice(user *cache.User, bot *telego.Bot, update *telego.Update) error {
	return nil
}

func handleEnterMinBonuses(user *cache.User, bot *telego.Bot, update *telego.Update) error {
	return nil
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
	case protobufs.UserState_EnterProductPool:
		err = handleEnterProductPool(user, bot, &update)
	case protobufs.UserState_EnterMinPrice:
		err = handleEnterMinPrice(user, bot, &update)
	case protobufs.UserState_EnterMinBonuses:
		err = handleEnterMinBonuses(user, bot, &update)
	}

	if err != nil {
		fmt.Println(err)
	}
}

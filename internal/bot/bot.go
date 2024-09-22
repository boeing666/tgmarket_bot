package bot

import (
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func StartMenu(bot *telego.Bot, update telego.Update) {
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			protobufs.ButtonAddProduct(),
			protobufs.ButtonProductList(),
		),
	)

	bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		GetWelcomeText(),
	).WithReplyMarkup(inlineKeyboard))
}

func Run(token string) error {
	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		return err
	}

	registerCallbacks()

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(StartMenu, th.CommandEqual("start"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Неизвестная команда, используйте /start",
		))
	}, th.AnyCommand())

	bh.Start()
	return nil
}

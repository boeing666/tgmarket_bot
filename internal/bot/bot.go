package bot

import (
	"context"
	"encoding/base64"
	"fmt"
	"tgmarket/internal/app"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/protobuf/proto"
)

func handleStartMenu(bot *telego.Bot, update telego.Update) {
	bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		welcomeText(),
	).WithReplyMarkup(protobufs.BuildMainMenu()))
}

func handleQueries(bot *telego.Bot, update telego.Update) {
	msgdate := uint64(update.CallbackQuery.Message.GetDate())

	app := app.GetApp()
	if msgdate < app.LaunchTime {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:    tu.ID(update.CallbackQuery.From.ID),
			Text:      "Сообщение устарело, /start чтобы начать работать с ботом.",
			MessageID: update.CallbackQuery.Message.GetMessageID(),
		})
		return
	}

	decode, err := base64.StdEncoding.DecodeString(update.CallbackQuery.Data)
	if err != nil {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:    tu.ID(update.CallbackQuery.From.ID),
			Text:      "Ошибка обработки сообщения, начните заново /start.",
			MessageID: update.CallbackQuery.Message.GetMessageID(),
		})
		return
	}

	var message protobufs.ButtonData
	err = proto.Unmarshal(decode, &message)
	if err != nil {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:    tu.ID(update.CallbackQuery.From.ID),
			Text:      "Ошибка обработки сообщения, начните заново /start.",
			MessageID: update.CallbackQuery.Message.GetMessageID(),
		})
		return
	}

	if callback, ok := buttonCallbacks[message.Id]; ok {
		ctx := callbackContext{bot, &update, message.Data}
		err := callback(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func handleMiddleware(bot *telego.Bot, update telego.Update, next th.Handler) {
	var userID int64
	if update.Message != nil {
		userID = update.Message.From.ID
	} else if update.CallbackQuery != nil {
		userID = update.CallbackQuery.From.ID
	}

	app := app.GetApp()
	ctx := update.Context()
	if userID != 0 {
		user := app.GetUser(userID)
		ctx = context.WithValue(ctx, "user", user)
	}

	update = update.WithContext(ctx)
	next(bot, update)
}

func Run(token string) error {
	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		return err
	}

	registerButtons()

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	bh.Use(handleMiddleware)
	bh.Handle(handleStartMenu, th.CommandEqual("start"))
	bh.Handle(handleUserStates, th.AnyMessage())
	bh.Handle(handleQueries, th.AnyCallbackQuery())

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Start()
	return nil
}

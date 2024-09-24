package bot

import (
	"context"
	"encoding/base64"
	"fmt"
	"tgmarket/internal/app"
	"tgmarket/internal/cache"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/protobuf/proto"
)

func handleStartMenu(bot *telego.Bot, update telego.Update) {
	ctx := update.Context()
	user := ctx.Value("user").(*cache.User)

	if user.ActiveMsgID != 0 {
		bot.DeleteMessage(tu.Delete(tu.ID(update.Message.Chat.ID), user.ActiveMsgID))
		user.ActiveMsgID = 0
	}

	res, err := bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		welcomeText(),
	).WithReplyMarkup(protobufs.BuildMainMenu()))

	if err == nil {
		user.ActiveMsgID = res.MessageID
	}
}

func handleQueries(ctx context.Context, bot *telego.Bot, query telego.CallbackQuery) {
	msgdate := uint64(query.Message.GetDate())

	app := app.GetApp()
	if msgdate < app.LaunchTime {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:    tu.ID(query.From.ID),
			Text:      "Сообщение устарело, /start чтобы начать работать с ботом.",
			MessageID: query.Message.GetMessageID(),
		})
		return
	}

	decode, err := base64.StdEncoding.DecodeString(query.Data)
	if err != nil {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:    tu.ID(query.From.ID),
			Text:      "Ошибка обработки сообщения, начните заново /start.",
			MessageID: query.Message.GetMessageID(),
		})
		return
	}

	var message protobufs.ButtonData
	err = proto.Unmarshal(decode, &message)
	if err != nil {
		bot.EditMessageText(&telego.EditMessageTextParams{
			ChatID:    tu.ID(query.From.ID),
			Text:      "Ошибка обработки сообщения, начните заново /start.",
			MessageID: query.Message.GetMessageID(),
		})
		return
	}

	if callback, ok := buttonCallbacks[message.Id]; ok {
		data := callbackData{
			baseContext: baseContext{
				bot:  bot,
				user: ctx.Value("user").(*cache.User),
			},
			data: message.GetData(),
		}
		if err := callback(&data); err != nil {
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

	ctx := update.Context()
	if userID != 0 {
		user, err := userscache.getUser(userID)
		if err != nil {
			fmt.Printf("Can't GetUser %s\n", err.Error())
			return
		}
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

	initUsersCache()

	registerStates()
	registerButtons()

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	bh.Use(handleMiddleware)
	bh.Handle(handleStartMenu, th.CommandEqual("start"))
	bh.HandleMessageCtx(handleUserStates)
	bh.HandleCallbackQueryCtx(handleQueries)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Start()
	return nil
}

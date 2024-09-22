package bot

import (
	"fmt"
	"tgmarket/internal/cache"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
)

func handleEnterProductURL(user *cache.User, bot *telego.Bot, update telego.Update) error {
	market := detectMarketplace(update.Message.Text)

	return nil
}

func handleEnterProductName(user *cache.User, bot *telego.Bot, update telego.Update) error {
	return nil
}

func handleEnterProductPool(user *cache.User, bot *telego.Bot, update telego.Update) error {
	return nil
}

func handleEnterMinPrice(user *cache.User, bot *telego.Bot, update telego.Update) error {
	return nil
}

func handleEnterMinBonuses(user *cache.User, bot *telego.Bot, update telego.Update) error {
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
		err = handleEnterProductURL(user, bot, update)
	case protobufs.UserState_EnterProductName:
		err = handleEnterProductName(user, bot, update)
	case protobufs.UserState_EnterProductPool:
		err = handleEnterProductPool(user, bot, update)
	case protobufs.UserState_EnterMinPrice:
		err = handleEnterMinPrice(user, bot, update)
	case protobufs.UserState_EnterMinBonuses:
		err = handleEnterMinBonuses(user, bot, update)
	}

	if err != nil {
		fmt.Println(err)
	}
}

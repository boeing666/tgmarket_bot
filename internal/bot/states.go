package bot

import (
	"fmt"
	"tgmarket/internal/protobufs"
	"tgmarket/internal/user"

	"github.com/mymmrac/telego"
)

func handleEnterProductURL(user *user.Cache, bot *telego.Bot, update telego.Update) error {
	return nil
}
func handleEnterProductName(user *user.Cache, bot *telego.Bot, update telego.Update) error {
	return nil
}
func handleEnterProductPool(user *user.Cache, bot *telego.Bot, update telego.Update) error {
	return nil
}
func handleEnterMinPrice(user *user.Cache, bot *telego.Bot, update telego.Update) error {
	return nil
}
func handleEnterMinBonuses(user *user.Cache, bot *telego.Bot, update telego.Update) error {
	return nil
}

func handleUserStates(bot *telego.Bot, update telego.Update) {
	ctx := update.Context()
	user := ctx.Value("user").(*user.Cache)

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

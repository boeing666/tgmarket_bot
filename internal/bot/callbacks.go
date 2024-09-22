package bot

import (
	"tgmarket/internal/protobufs"

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

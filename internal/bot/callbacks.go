package bot

import (
	"tgmarket/internal/protobufs"

	tu "github.com/mymmrac/telego/telegoutil"
)

func callbackAddProduct(ctx callbackContext) error {
	ctx.EditMessageTextWithKeyboard(addNewProductText(),
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
	return nil
}

func callbackChangeMenu(ctx callbackContext) error {
	return nil
}

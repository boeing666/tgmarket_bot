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
	if err := ctx.Unmarshal(&msg); err != nil {
		return err
	}

	return showProductInfo(msg.Id, ctx.Bot, ctx.GetUser())
}

func callbackListOfProducts(ctx callbackContext) error {
	user := ctx.GetUser()

	if ctx.Data != nil {
		var msg protobufs.ChangePage
		if err := ctx.Unmarshal(&msg); err != nil {
			return err
		}
		user.LastPage = int(*msg.Newpage)
	}

	var rows [][]telego.InlineKeyboardButton
	products, menuInfo := buildPage(user.LastPage, user.Products)

	header := buildMenuHeader(menuInfo)
	if header != nil {
		rows = append(rows, buildMenuHeader(menuInfo))
	}

	for _, id := range products {
		product := user.Products[id]
		rows = append(rows,
			protobufs.CreateRowButton(
				fmt.Sprintf("%s (%s)", product.Name, getShopName(protobufs.Shops(product.ShopID))),
				protobufs.ButtonID_ProductInfo,
				&protobufs.ProdcutData{Id: id},
			),
		)
	}

	rows = append(rows,
		buildNavigation(menuInfo, protobufs.ButtonID_ListOfProducts),
		tu.InlineKeyboardRow(protobufs.ButtonBack(protobufs.ButtonID_MainMenu, nil)),
	)

	user.LastPage = menuInfo.page
	_, err := ctx.EditLastMessage(listOfProductsText(), &telego.InlineKeyboardMarkup{InlineKeyboard: rows})
	return err
}

func callbackSetMinPrice(ctx callbackContext) error {
	user := ctx.GetUser()
	user.State = protobufs.UserState_EnterMinPrice
	_, err := ctx.EditLastMessage(enterMinProductPriceText(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(protobufs.ButtonCancelProduct(user.ActiveProductID)),
		),
	)
	return err
}

func callbackSetMinBonuses(ctx callbackContext) error {
	user := ctx.GetUser()
	user.State = protobufs.UserState_EnterMinBonuses
	_, err := ctx.EditLastMessage(enterMinProductBonuses(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(protobufs.ButtonCancelProduct(user.ActiveProductID)),
		),
	)
	return err
}

func callbackChangeProductName(ctx callbackContext) error {
	user := ctx.GetUser()
	user.State = protobufs.UserState_EnterProductName
	_, err := ctx.EditLastMessage(enterProductNameText(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(protobufs.ButtonCancelProduct(user.ActiveProductID)),
		),
	)
	return err
}

func callbackDeleteProduct(ctx callbackContext) error {
	user := ctx.GetUser()
	product, found := user.Products[user.ActiveProductID]
	if !found {
		return fmt.Errorf("DeleteProduct id %d user %d", product.ID, user.TelegramID)
	}

	if err := user.RemoveProduct(product.ID); err != nil {
		return err
	}

	if user.LastPage != -1 {
		callbackListOfProducts(ctx)
	} else {
		callbackMainMenu(ctx)
	}

	return nil
}

func callbackMainMenu(ctx callbackContext) error {
	user := ctx.GetUser()
	user.State = protobufs.UserState_None
	_, err := ctx.EditLastMessage(welcomeText(), protobufs.BuildMainMenu())
	return err
}

func callbackChangeMenu(ctx callbackContext) error {
	var msg protobufs.ButtonData
	if err := ctx.Unmarshal(&msg); err != nil {
		return err
	}
	ctx.Data = msg.Data
	return buttonCallbacks[msg.Id](ctx)
}

func callbackNothing(ctx callbackContext) error {
	err := ctx.Bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
		CallbackQueryID: ctx.Update.CallbackQuery.ID,
		Text:            "Тут ничего нет!",
	})
	return err
}

func showProductInfo(productID int64, bot *telego.Bot, user *cache.User) error {
	product, found := user.Products[productID]
	if !found {
		return fmt.Errorf("Not found product id %d user %d", productID, user.TelegramID)
	}

	user.ActiveProductID = productID

	createDate := product.CreatedAt.Format("2006-01-02 15:04:05")
	updateDate := product.UpdatedAt.Format("2006-01-02 15:04:05")

	var backbutton []telego.InlineKeyboardButton
	if user.LastPage != -1 {
		backbutton = tu.InlineKeyboardRow(protobufs.ButtonBack(protobufs.ButtonID_ListOfProducts, nil))
	} else {
		backbutton = tu.InlineKeyboardRow(protobufs.ButtonBack(protobufs.ButtonID_MainMenu, nil))
	}

	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(protobufs.ButtonSetMinimalPrice()),
		tu.InlineKeyboardRow(protobufs.ButtonSetMinimalBonuses()),
		tu.InlineKeyboardRow(protobufs.ButtonSetProductName()),
		tu.InlineKeyboardRow(protobufs.ButtonDeleteProduct()),
		backbutton,
		tu.InlineKeyboardRow(protobufs.ButtonMainMenu()),
	)

	text, entities := tu.MessageEntities(
		tu.Entity("📦 Товар: "), tu.Entity(fmt.Sprintf("%s\n", product.Name)).TextLink(product.URL),
		tu.Entity(fmt.Sprintf("🛒 Магазин: %s\n", getShopName(protobufs.Shops(product.ShopID)))),
		tu.Entity(fmt.Sprintf("💰 Цена: %d руб\n", product.Price)),
		tu.Entity(fmt.Sprintf("💵 Мин цена: %d\n", product.MinPrice)),
		tu.Entity(fmt.Sprintf("❇️ Мин кол-во бонусов: %d\n", product.MinBonuses)),
		tu.Entity(fmt.Sprintf("🗓️ Добавлен: %s\n", createDate)),
		tu.Entity(fmt.Sprintf("🗓️ Обновлен: %s\n", updateDate)),
	)

	bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:             tu.ID(user.TelegramID),
		Text:               text,
		MessageID:          user.ActiveMsgID,
		Entities:           entities,
		ReplyMarkup:        keyboard,
		LinkPreviewOptions: &telego.LinkPreviewOptions{IsDisabled: true},
	})

	return nil
}

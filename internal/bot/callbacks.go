package bot

import (
	"fmt"
	"tgmarket/internal/cache"
	"tgmarket/internal/models"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/protobuf/proto"
)

type callbackHandler func(data messageContext) error

var buttonCallbacks map[protobufs.ButtonID]callbackHandler

func callbackAddProduct(data messageContext) error {
	button := protobufs.ButtonCancel(protobufs.ButtonID_MainMenu, nil)
	data.EditLastMessage(addNewProductText(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(button),
		),
	)
	data.GetUser().State = protobufs.UserState_EnterProductURL
	return nil
}

func callbackSearchByName(data messageContext) error {
	button := protobufs.ButtonCancel(protobufs.ButtonID_ListOfProducts, nil)
	data.EditLastMessage(enterProductNameForFilter(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(button),
		),
	)
	data.GetUser().State = protobufs.UserState_EnterPartialProductName
	return nil
}

func callbackRemoveFilterByName(data messageContext) error {
	user := data.GetUser()
	user.FilterName = ""
	user.FiltredProducts = make(map[int64]*models.Product)
	return callbackListOfProducts(data)
}

func callbackProductInfo(data messageContext) error {
	var msg protobufs.ProdcutData
	if err := proto.Unmarshal(data.GetCallbackData(), &msg); err != nil {
		return err
	}

	return showProductInfo(msg.GetId(), data.GetBot(), data.GetUser())
}

func callbackListOfProducts(data messageContext) error {
	user := data.GetUser()
	user.State = protobufs.UserState_None

	if len(user.Products) == 0 {
		return data.GetBot().AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
			CallbackQueryID: data.GetCallbackQueryID(),
			Text:            "Ваш список товаров пуст, добавьте что-нибудь!",
		})
	}

	if data.GetCallbackData() != nil {
		var msg protobufs.ChangePage
		if err := proto.Unmarshal(data.GetCallbackData(), &msg); err != nil {
			return err
		}
		user.LastPage = int(*msg.Newpage)
	}

	var rows [][]telego.InlineKeyboardButton

	var listOfProducts map[int64]*models.Product
	if len(user.FilterName) != 0 {
		listOfProducts = user.FiltredProducts
	} else {
		listOfProducts = user.Products
	}

	products, menuInfo := buildPage(user.LastPage, listOfProducts)

	header := buildMenuHeader(user, menuInfo)
	if header != nil {
		rows = append(rows, header...)
	}

	for _, id := range products {
		product := listOfProducts[id]
		rows = append(rows,
			protobufs.CreateRowButton(
				fmt.Sprintf("%s (%s)", product.Name, getShopName(protobufs.Shops(product.ShopID))),
				protobufs.ButtonID_ProductInfo,
				&protobufs.ProdcutData{Id: id},
			),
		)
	}

	navigation := buildNavigation(menuInfo, protobufs.ButtonID_ListOfProducts)
	if navigation != nil {
		rows = append(rows,
			navigation,
		)
	}

	rows = append(rows,
		tu.InlineKeyboardRow(protobufs.ButtonBack(protobufs.ButtonID_MainMenu, nil)),
	)

	user.LastPage = menuInfo.page
	_, err := data.EditLastMessage(listOfProductsText(), &telego.InlineKeyboardMarkup{InlineKeyboard: rows})
	return err
}

func callbackSetMinPrice(data messageContext) error {
	user := data.GetUser()
	user.State = protobufs.UserState_EnterMinPrice
	button := protobufs.ButtonCancel(
		protobufs.ButtonID_ProductInfo,
		&protobufs.ProdcutData{Id: user.ActiveProductID},
	)
	_, err := data.EditLastMessage(enterMinProductPriceText(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(button),
		),
	)
	return err
}

func callbackSetMinBonuses(data messageContext) error {
	user := data.GetUser()
	user.State = protobufs.UserState_EnterMinBonuses
	button := protobufs.ButtonCancel(
		protobufs.ButtonID_ProductInfo,
		&protobufs.ProdcutData{Id: user.ActiveProductID},
	)
	_, err := data.EditLastMessage(enterMinProductBonuses(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(button),
		),
	)
	return err
}

func callbackChangeProductName(data messageContext) error {
	user := data.GetUser()
	user.State = protobufs.UserState_EnterProductName
	button := protobufs.ButtonCancel(
		protobufs.ButtonID_ProductInfo,
		&protobufs.ProdcutData{Id: user.ActiveProductID},
	)
	_, err := data.EditLastMessage(enterProductNameText(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(button),
		),
	)
	return err
}

func callbackDeleteProduct(data messageContext) error {
	user := data.GetUser()
	product, found := user.Products[user.ActiveProductID]
	if !found {
		return fmt.Errorf("DeleteProduct id %d user %d", product.ID, user.TelegramID)
	}

	if err := user.RemoveProduct(product.ID); err != nil {
		return err
	}

	if len(user.Products) == 0 || user.LastPage == -1 {
		return callbackMainMenu(data)
	}

	return callbackListOfProducts(data)
}

func callbackMainMenu(data messageContext) error {
	user := data.GetUser()
	user.State = protobufs.UserState_None
	user.LastPage = -1
	_, err := data.EditLastMessage(welcomeText(), protobufs.BuildMainMenu())
	return err
}

func callbackChangeMenu(data messageContext) error {
	var msg protobufs.ButtonData
	if err := proto.Unmarshal(data.GetCallbackData(), &msg); err != nil {
		return err
	}
	data.SetCallbackData(msg.Data)
	return buttonCallbacks[msg.Id](data)
}

func callbackNothing(data messageContext) error {
	return data.GetBot().AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{
		CallbackQueryID: data.GetCallbackQueryID(),
		Text:            "Тут ничего нет!",
	})
}

func showProductInfo(productID int64, bot *telego.Bot, user *cache.User) error {
	user.State = protobufs.UserState_None

	product, found := user.Products[productID]
	if !found {
		return fmt.Errorf("Not found product id %d user %d", productID, user.TelegramID)
	}

	user.ActiveProductID = productID

	createDate := product.CreatedAt.Format("2006-01-02 15:04:05")
	updateDate := product.UpdatedAt.Format("2006-01-02 15:04:05")

	keyboard := [][]telego.InlineKeyboardButton{
		tu.InlineKeyboardRow(protobufs.ButtonSetMinimalPrice()),
		tu.InlineKeyboardRow(protobufs.ButtonSetMinimalBonuses()),
		tu.InlineKeyboardRow(protobufs.ButtonSetProductName()),
		tu.InlineKeyboardRow(protobufs.ButtonDeleteProduct()),
	}

	if user.LastPage != -1 {
		keyboard = append(keyboard, tu.InlineKeyboardRow(protobufs.ButtonBack(protobufs.ButtonID_ListOfProducts, nil)))
	}
	keyboard = append(keyboard, tu.InlineKeyboardRow(protobufs.ButtonMainMenu()))

	text, entities := tu.MessageEntities(
		tu.Entity("📦 Товар: "), tu.Entity(fmt.Sprintf("%s\n", product.Name)).TextLink(product.URL),
		tu.Entity(fmt.Sprintf("🛒 Магазин: %s\n", getShopName(protobufs.Shops(product.ShopID)))),
		tu.Entity(fmt.Sprintf("💰 Текущая цена: %d ₽\n", product.Price)),
		tu.Entity(fmt.Sprintf("💰 Сумма бонусов: %d\n", product.Bonus)),
		tu.Entity(fmt.Sprintf("💵 Мин цена: %d\n", product.MinPrice)),
		tu.Entity(fmt.Sprintf("❇️ Мин кол-во бонусов: %d\n", product.MinBonuses)),
		tu.Entity(fmt.Sprintf("🗓️ Добавлен: %s\n", createDate)),
		tu.Entity(fmt.Sprintf("🗓️ Обновлен: %s\n", updateDate)),
	)

	_, err := bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:             tu.ID(user.TelegramID),
		Text:               text,
		MessageID:          user.ActiveMsgID,
		Entities:           entities,
		ReplyMarkup:        &telego.InlineKeyboardMarkup{InlineKeyboard: keyboard},
		LinkPreviewOptions: &telego.LinkPreviewOptions{IsDisabled: true},
	})

	return err
}

func registerButtons() {
	buttonCallbacks = make(map[protobufs.ButtonID]callbackHandler)

	buttonCallbacks[protobufs.ButtonID_AddProduct] = callbackAddProduct
	buttonCallbacks[protobufs.ButtonID_ListOfProducts] = callbackListOfProducts
	buttonCallbacks[protobufs.ButtonID_SetMinPrice] = callbackSetMinPrice
	buttonCallbacks[protobufs.ButtonID_SetMinBonuses] = callbackSetMinBonuses
	buttonCallbacks[protobufs.ButtonID_ChangeProductName] = callbackChangeProductName
	buttonCallbacks[protobufs.ButtonID_DeleteProduct] = callbackDeleteProduct
	buttonCallbacks[protobufs.ButtonID_ChangeMenu] = callbackChangeMenu
	buttonCallbacks[protobufs.ButtonID_MainMenu] = callbackMainMenu
	buttonCallbacks[protobufs.ButtonID_ProductInfo] = callbackProductInfo
	buttonCallbacks[protobufs.ButtonID_SearchByName] = callbackSearchByName
	buttonCallbacks[protobufs.ButtonID_RemoveFilterByName] = callbackRemoveFilterByName
	buttonCallbacks[protobufs.ButtonID_Nothing] = callbackNothing
}

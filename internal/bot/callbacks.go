package bot

import (
	"fmt"
	"tgmarket/internal/cache"
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/protobuf/proto"
)

type callbackHandler func(data messageContext) error

var buttonCallbacks map[protobufs.ButtonID]callbackHandler

func callbackAddProduct(data messageContext) error {
	data.EditLastMessage(addNewProductText(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(protobufs.ButtonCancel()),
		),
	)
	data.GetUser().State = protobufs.UserState_EnterProductURL
	return nil
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

	if data.GetCallbackData() != nil {
		var msg protobufs.ChangePage
		if err := proto.Unmarshal(data.GetCallbackData(), &msg); err != nil {
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
	_, err := data.EditLastMessage(listOfProductsText(), &telego.InlineKeyboardMarkup{InlineKeyboard: rows})
	return err
}

func callbackSetMinPrice(data messageContext) error {
	user := data.GetUser()
	user.State = protobufs.UserState_EnterMinPrice
	_, err := data.EditLastMessage(enterMinProductPriceText(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(protobufs.ButtonCancelProduct(user.ActiveProductID)),
		),
	)
	return err
}

func callbackSetMinBonuses(data messageContext) error {
	user := data.GetUser()
	user.State = protobufs.UserState_EnterMinBonuses
	_, err := data.EditLastMessage(enterMinProductBonuses(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(protobufs.ButtonCancelProduct(user.ActiveProductID)),
		),
	)
	return err
}

func callbackChangeProductName(data messageContext) error {
	user := data.GetUser()
	user.State = protobufs.UserState_EnterProductName
	_, err := data.EditLastMessage(enterProductNameText(),
		tu.InlineKeyboard(
			tu.InlineKeyboardRow(protobufs.ButtonCancelProduct(user.ActiveProductID)),
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

	if user.LastPage != -1 {
		return callbackListOfProducts(data)
	} else {
		return callbackMainMenu(data)
	}
}

func callbackMainMenu(data messageContext) error {
	data.GetUser().State = protobufs.UserState_None
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

	_, err := bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:             tu.ID(user.TelegramID),
		Text:               text,
		MessageID:          user.ActiveMsgID,
		Entities:           entities,
		ReplyMarkup:        keyboard,
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
	buttonCallbacks[protobufs.ButtonID_Nothing] = callbackNothing
}

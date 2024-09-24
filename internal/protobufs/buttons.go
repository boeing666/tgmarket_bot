package protobufs

import (
	"encoding/base64"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/protobuf/proto"
)

func BuildMainMenu() *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			ButtonAddProduct(),
			ButtonProductList(),
		),
	)
}

func ButtonAddProduct() telego.InlineKeyboardButton {
	return CreateButton("🆕✨Добавить Товар", ButtonID_AddProduct, nil)
}

func ButtonProductList() telego.InlineKeyboardButton {
	return CreateButton("🗂️📦 Список Товаров", ButtonID_ListOfProducts, nil)
}

func ButtonMainMenu() telego.InlineKeyboardButton {
	return CreateButton("↩️ На главную", ButtonID_MainMenu, nil)
}

func ButtonSetMinimalPrice() telego.InlineKeyboardButton {
	return CreateButton("💵 Установить минимальную цену", ButtonID_SetMinPrice, nil)
}

func ButtonSetMinimalBonuses() telego.InlineKeyboardButton {
	return CreateButton("❇️Установить минимальные бонусы", ButtonID_SetMinBonuses, nil)
}

func ButtonSetProductName() telego.InlineKeyboardButton {
	return CreateButton("✏️ Изменить имя товара", ButtonID_ChangeProductName, nil)
}

func ButtonDeleteProduct() telego.InlineKeyboardButton {
	return CreateButton("🗑️ Удалить Товар", ButtonID_DeleteProduct, nil)
}

func ButtonCancel() telego.InlineKeyboardButton {
	return CreateButton("❌ Отменить ввод", ButtonID_MainMenu, nil)
}

func ButtonCancelProduct(id int64) telego.InlineKeyboardButton {
	return CreateButton("❌ Отменить ввод", ButtonID_ProductInfo, &ProdcutData{Id: id})
}

func ButtonBack(newmenu ButtonID, msg proto.Message) telego.InlineKeyboardButton {
	var bytes []byte
	if msg != nil {
		bytes, _ = proto.Marshal(msg)
	}

	data := ButtonData{Id: newmenu, Data: bytes}
	return CreateButton("⬅️ Назад", ButtonID_ChangeMenu, &data)
}

func CreateButton(name string, btnID ButtonID, msg proto.Message) telego.InlineKeyboardButton {
	var bytes []byte
	var err error

	if msg != nil {
		bytes, err = proto.Marshal(msg)
		if err != nil {
			fmt.Print(err)
		}
	}

	button := ButtonData{Id: btnID, Data: bytes}
	bytes, _ = proto.Marshal(&button)

	encoded := base64.StdEncoding.EncodeToString(bytes)
	return tu.InlineKeyboardButton(name).WithCallbackData(encoded)
}

func CreateRowButton(name string, messageID ButtonID, msg proto.Message) []telego.InlineKeyboardButton {
	return tu.InlineKeyboardRow(CreateButton(name, messageID, msg))
}

package protobufs

import (
	"encoding/base64"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/protobuf/proto"
)

func ButtonAddProduct() telego.InlineKeyboardButton {
	return CreateButton("🆕✨Добавить Товар", ButtonID_AddProduct, nil)
}

func ButtonProductList() telego.InlineKeyboardButton {
	return CreateButton("🗂️📦 Список Товаров", ButtonID_ListOfProducts, nil)
}

func ButtonMainMenu() telego.InlineKeyboardButton {
	data := ButtonData{Id: ButtonID_MainMenu}
	return CreateButton("↩️ На главную", ButtonID_ChangeMenu, &data)
}

func ButtonCancel() telego.InlineKeyboardButton {
	data := ButtonData{Id: ButtonID_MainMenu}
	return CreateButton("❌ Отменить ввод", ButtonID_MainMenu, &data)
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

	if msg != nil {
		bytes, _ = proto.Marshal(msg)
	}

	button := ButtonData{Id: btnID, Data: bytes}
	bytes, _ = proto.Marshal(&button)

	encoded := base64.StdEncoding.EncodeToString(bytes)
	return tu.InlineKeyboardButton(name).WithCallbackData(encoded)
}

func CreateRowButton(name string, messageID ButtonID, msg proto.Message) []telego.InlineKeyboardButton {
	return tu.InlineKeyboardRow(
		CreateButton(name, messageID, msg),
	)
}

package protobufs

import (
	"encoding/base64"

	"github.com/gogo/protobuf/proto"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ButtonAddProduct() telego.InlineKeyboardButton {
	return CreateButton("🆕✨Добавить Товар", ButtonID_AddProduct, nil)
}

func ButtonProductList() telego.InlineKeyboardButton {
	return CreateButton("🗂️📦 Список Товаров", ButtonID_ListOfProducts, nil)
}

func ButtonMainMenu() telego.InlineKeyboardButton {
	data := ChangeMenuData{Newmenu: ButtonID_MainMenu}
	return CreateButton("↩️ На главную", ButtonID_ChangeMenu, &data)
}

func ButtonBack(newmenu ButtonID, msg proto.Message) telego.InlineKeyboardButton {
	var bytes []byte
	if msg != nil {
		bytes, _ = proto.Marshal(msg)
	}

	data := ChangeMenuData{Newmenu: newmenu, Data: bytes}
	return CreateButton("⬅️ Назад", ButtonID_ChangeMenu, &data)
}

func CreateButton(name string, btnID ButtonID, msg proto.Message) telego.InlineKeyboardButton {
	var bytes []byte

	if msg != nil {
		bytes, _ = proto.Marshal(msg)
	}

	fullmsg, _ := proto.Marshal(&ButtonData{Id: btnID, Data: bytes})
	encoded := base64.StdEncoding.EncodeToString(fullmsg)

	return tu.InlineKeyboardButton(name).WithCallbackData(encoded)
}

func CreateRowButton(name string, messageID ButtonID, msg proto.Message) []telego.InlineKeyboardButton {
	return tu.InlineKeyboardRow(
		CreateButton(name, messageID, msg),
	)
}

package bot

import (
	"tgmarket/internal/protobufs"

	"github.com/mymmrac/telego"
)

type callbackContext struct {
	Bot    *telego.Bot
	Update *telego.Update
	Data   []byte
}

type queryHandler func(ctx callbackContext) error

var buttonCallbacks map[protobufs.ButtonID]queryHandler

func registerCallbacks() {
	buttonCallbacks = make(map[protobufs.ButtonID]queryHandler)

	buttonCallbacks[protobufs.ButtonID_AddProduct] = callbackAddProduct
	buttonCallbacks[protobufs.ButtonID_ListOfProducts] = callbackListOfProducts
}

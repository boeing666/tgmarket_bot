package bot

import (
	"tgmarket/internal/cache"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type baseContext struct {
	bot  *telego.Bot
	user *cache.User
}

func (b *baseContext) EditLastMessageText(text string) (*telego.Message, error) {
	return b.bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:    tu.ID(b.user.TelegramID),
		Text:      text,
		MessageID: b.user.ActiveMsgID,
	})
}

func (b *baseContext) EditLastMessage(text string, keyboard *telego.InlineKeyboardMarkup) (*telego.Message, error) {
	return b.bot.EditMessageText(&telego.EditMessageTextParams{
		ChatID:      tu.ID(b.user.TelegramID),
		Text:        text,
		MessageID:   b.user.ActiveMsgID,
		ReplyMarkup: keyboard,
	})
}

type messageContext interface {
	GetBot() *telego.Bot
	GetUser() *cache.User
	SetCallbackData(data []byte)
	GetCallbackData() []byte
	GetCallbackQueryID() string
	GetMessageID() int
	GetMessageText() string
	EditLastMessageText(text string) (*telego.Message, error)
	EditLastMessage(text string, keyboard *telego.InlineKeyboardMarkup) (*telego.Message, error)
}

type messageData struct {
	baseContext

	messageid   int
	messagetext string
}

func (m *messageData) GetBot() *telego.Bot         { return m.bot }
func (m *messageData) GetUser() *cache.User        { return m.user }
func (m *messageData) SetCallbackData(data []byte) {}
func (m *messageData) GetCallbackData() []byte     { return nil }
func (m *messageData) GetMessageID() int           { return m.messageid }
func (m *messageData) GetMessageText() string      { return m.messagetext }
func (m *messageData) GetCallbackQueryID() string  { return "" }

type callbackData struct {
	baseContext

	data    []byte
	queryid string
}

func (c *callbackData) GetBot() *telego.Bot         { return c.bot }
func (c *callbackData) GetUser() *cache.User        { return c.user }
func (c *callbackData) SetCallbackData(data []byte) { c.data = data }
func (c *callbackData) GetCallbackData() []byte     { return c.data }
func (c *callbackData) GetMessageID() int           { return 0 }
func (c *callbackData) GetMessageText() string      { return "" }
func (c *callbackData) GetCallbackQueryID() string  { return c.queryid }

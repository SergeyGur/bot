package commands

import (
	"log"

	"github.com/SegeyGur/bot/internal/sevice/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update, productService *product.Service) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic : %v", panicValue)
		}
	}()

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message, productService)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}

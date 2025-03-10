package commands

import (
	"github.com/SegeyGur/bot/internal/sevice/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMessage := "Here all products : \n\n"

	products := productService.List()
	for _, p := range products {
		outputMessage += p.Title
		outputMessage += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)

	c.bot.Send(msg)
}

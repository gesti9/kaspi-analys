package service

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	token        = "6553780269:AAGKRvVeV7cswTqcjEErQKbBfdU6t6cYE-Y"
	paymentToken = os.Getenv("5420394252:TEST:543267")
)

func Pay(id int) {
	bot, _ := tgbotapi.NewBotAPI(token)
	invoice := tgbotapi.NewInvoice(
		int64(id),
		"Оплата за подписку!",
		"Платеж на сумму 4990₸",
		"custom_payload",
		paymentToken, // Токен для создания платежа
		"start_param",
		"KZT",
		&[]tgbotapi.LabeledPrice{{Label: "KZT", Amount: 499000}},
	)
	invoice.ProviderToken = "5420394252:TEST:543267"

	log.Println("Before sending invoice")
	_, _ = bot.Send(invoice)

	log.Println("After sending invoice")
}

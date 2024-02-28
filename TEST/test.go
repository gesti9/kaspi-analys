package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	var (
		paymentToken = os.Getenv("5420394252:TEST:543267")
	)

	bot, err := tgbotapi.NewBotAPI("6775506137:AAGh1_Xs7s86dpIiop2jW1bpxMKknh5AEJQ")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			invoice := tgbotapi.NewInvoice(
				update.Message.Chat.ID,
				"Оплата за подписку!",
				"Платеж на сумму 4990₸",
				"custom_payload",
				paymentToken, // Токен для создания платежа
				"start_param",
				"KZT",
				&[]tgbotapi.LabeledPrice{{Label: "KZT", Amount: 500000}},
			)
			invoice.ProviderToken = "5420394252:TEST:543267"

			log.Println("Before sending invoice")
			_, err = bot.Send(invoice)
			if err != nil {
				log.Println("Error sending invoice:", err)
			}
			log.Println("After sending invoice")

			msgText := ""
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			bot.Send(msg)

			// Дополнительный код по обработке инвойса, если нужно
		}
	}
}

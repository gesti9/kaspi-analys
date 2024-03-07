package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"work/data"
	"work/logs"
	"work/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserState struct {
	CurrentState string
	PrevState    string
}

var (
	bot             *tgbotapi.BotAPI
	userStates      = make(map[int64]*UserState)
	userStatesMutex sync.Mutex
	mainMenu        = tgbotapi.NewReplyKeyboard(

		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Купить подписку!"),
			tgbotapi.NewKeyboardButton("Поддержка!"),
		),
	)
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6553780269:AAGKRvVeV7cswTqcjEErQKbBfdU6t6cYE-Y")
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			count, _ := strconv.Atoi(data.ReadFromFile("data/users/" + strconv.Itoa(int(update.Message.Chat.ID)) + ".txt"))
			switch update.Message.Text {
			case "/start":
				logs.Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Жду ссылку..")
				msg.ReplyToMessageID = update.Message.MessageID
				if update.Message != nil && update.Message.Contact != nil {
					// Теперь можно безопасно использовать PhoneNumber
					fmt.Println(update.Message.Contact.PhoneNumber)
				} else {
					fmt.Println("Контакт не предоставлен или объект Contact равен nil")
				}
				msg.ReplyMarkup = mainMenu
				bot.Send(msg)
			case "Купить подписку!":
				logs.Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				// service.Pay(int(update.Message.Chat.ID))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, `	Kaspi оплата 4990₸`)
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("Оплата", "https://pay.kaspi.kz/pay/jxrd4qnx"),
					),
				)
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, `После оплаты напишите в поддержку!`)
				bot.Send(msg)
			case "/payments":
				logs.Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				// service.Pay(int(update.Message.Chat.ID))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, `	Kaspi оплата 4990₸`)
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("Оплата", "https://pay.kaspi.kz/pay/jxrd4qnx"),
					),
				)
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, `После оплаты напишите в поддержку!`)
				bot.Send(msg)
			case "Поддержка!":
				logs.Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Переходи👇`)
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("Поддержка!", "https://t.me/erasild"),
					),
				)
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			case "/helper":
				logs.Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Переходи👇`)
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("Поддержка!", "https://t.me/erasild"),
					),
				)
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			default:
				logs.Log("@" + update.Message.From.UserName + "  " + "ИМЯ: " + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "  " + "ID: " + strconv.Itoa(int(update.Message.Chat.ID)) + "  " + update.Message.Text + "\n")
				if service.IsValidURL(update.Message.Text) {
					fmt.Printf("%s - это валидная ссылка\n", (update.Message.Text))
					result, _ := service.Output(update.Message.Text)
					num, _ := strconv.Atoi(result)

					if data.ReadFromFile("data/users/"+strconv.Itoa(int(update.Message.Chat.ID))+".txt") == "10" {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Закончился пробный период, для продолжения купите подписку!")
						msg.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg)
						fmt.Println(data.ReadFromFile("data/users/" + strconv.Itoa(int(update.Message.Chat.ID)) + ".txt"))
					} else if num == 0 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "У товара 0 продаж!")
						msg.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg)
					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Идет обработка..")
						bot.Send(msg)
						mes := (float64(num) / float64(365)) * 30
						day := float64(mes) / float64(30)
						formatted := fmt.Sprintf("%.1f", day)
						fmt.Println(formatted)
						price, _ := service.Price(update.Message.Text)
						moneyM := price * 30
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, `Общее количество продаж: `+result+` шт.`+"\n"+
							`В месяц: `+strconv.Itoa(int(mes))+` шт.`+"\n"+`Продажи в день: `+formatted+` шт.`+
							"\n"+`Выручка за месяц: `+strconv.Itoa(moneyM)+` ₸`)
						msg.ReplyToMessageID = update.Message.MessageID

						bot.Send(msg)

						count++
						data.UserData(update.Message.From.ID, count)
					}

					// Ваш код для загрузки и анализа страницы
				} else {
					fmt.Printf("%s - не является валидной ссылкой\n", (update.Message.Text))
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нужна ссылка с каспи!!!!")
					msg.ReplyToMessageID = update.Message.MessageID

					bot.Send(msg)
				}

			}

		}
	}
}

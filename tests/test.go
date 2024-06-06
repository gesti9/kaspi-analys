package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func Price(url string) (int, error) {
	// Создаем контекст
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Настройка времени ожидания
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var res string

	// Выполняем задачи
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second), // Ждем 2 секунды для загрузки страницы
		chromedp.Text(`.item__price-once`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		return 0, err
	}

	fmt.Println("Текст из элемента item__price-once:", res)

	// Извлекаем только цифры с использованием регулярного выражения
	re := regexp.MustCompile(`\d+`)
	digits := re.FindAllString(res, -1)

	// Объединяем извлеченные цифры в одну строку
	resultString := ""
	for _, digit := range digits {
		resultString += digit
	}

	// Преобразуем строку в число
	price, err := strconv.Atoi(resultString)
	if err != nil {
		return 0, err
	}

	fmt.Println("Извлеченные цифры:", price)

	return price, nil
}

func main() {
	url := "https://kaspi.kz/shop/p/magnum-banan-ekvador-101349284/?c=750000000&utm_source=google&utm_medium=cpc&utm_campaign=shop_google_performance_max_car_goods&gclid=CjwKCAjwvIWzBhAlEiwAHHWgvUw7z22sNzQYgdx6ORJkdTYXNLSwsYvaY52s8N399ZGVwCP5TbAt6xoCaD8QAvD_BwE"
	price, err := Price(url)
	if err != nil {
		log.Fatal("Ошибка получения цены:", err)
	}
	fmt.Println("Цена:", price)
}

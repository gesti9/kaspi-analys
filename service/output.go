package service

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func Output(n string) (string, error) {
	url := n
	var result string

	// Отправка GET-запроса
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	// Выполнение запроса с http.DefaultClient
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("ошибка запроса. Статус: %d", resp.StatusCode)
	}

	// Используем goquery для парсинга HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании документа из ответа: %v", err)
	}

	// Используем регулярное выражение для поиска значения "reviewsCount"
	re := regexp.MustCompile(`"reviewsCount":(\d+)`)

	// Поиск по тексту страницы
	doc.Find("script").Each(func(index int, element *goquery.Selection) {
		scriptText := element.Text()
		match := re.FindStringSubmatch(scriptText)
		if len(match) == 2 {
			reviewsCount := match[1]
			fmt.Printf("Значение reviewsCount: %s\n", reviewsCount)

			result = reviewsCount
		}
	})

	return strconv.Itoa(int(Sum(result))), nil
}

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

// Функция для вычисления суммы (ваша логика)
func Sum(s string) float64 {
	num, _ := strconv.Atoi(s)
	return float64(57*num) / float64(43)
}

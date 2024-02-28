package service

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
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

// Функция для вычисления суммы (ваша логика)
func Sum(s string) float64 {
	num, _ := strconv.Atoi(s)
	return float64(57*num) / float64(43)
}

// func Output(n string) string {
// 	url := n
// 	var result string
// 	// sum := 0

// 	// Отправка GET-запроса
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Выполнение запроса с http.DefaultClient
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	// Проверка статуса ответа
// 	if resp.StatusCode != 200 {
// 		log.Printf("Ошибка запроса. Статус: %d", resp.StatusCode)
// 		return ""
// 	}

// 	// Используем goquery для парсинга HTML
// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		log.Print("Ошибка при создании документа из ответа:", err)
// 		return ""
// 	}

// 	// Используем регулярное выражение для поиска значения "reviewsCount"
// 	re := regexp.MustCompile(`"reviewsCount":(\d+)`)

// 	// Поиск по тексту страницы
// 	doc.Find("script").Each(func(index int, element *goquery.Selection) {
// 		scriptText := element.Text()
// 		match := re.FindStringSubmatch(scriptText)
// 		if len(match) == 2 {
// 			reviewsCount := match[1]
// 			fmt.Printf("Значение reviewsCount: %s\n", reviewsCount)

// 			result = reviewsCount // преобразование в стро
// 		}
// 	})
// 	// num, _ := strconv.Atoi(result)

// 	// sum := float64(57*num) / float64(43)
// 	return strconv.Itoa(int(Sum(result)))
// }

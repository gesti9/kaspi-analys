package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func extractPrices(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("ошибка при выполнении GET-запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("ошибка при чтении тела ответа: %v", err)
	}

	// Найти все совпадения с регулярным выражением
	re := regexp.MustCompile(`"price": "(\d+)"`)
	matches := re.FindAllStringSubmatch(string(body), -1)

	// Извлечь только цифры из совпадений
	var sum int
	for _, match := range matches {
		if len(match) >= 2 {
			price, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, fmt.Errorf("ошибка при конвертации строки в число: %v", err)
			}
			sum += price
		}
	}

	return sum, nil
}

func main() {
	url := "https://kaspi.kz/shop/p/nb-f80-chernyi-110855908/?c=750000000&utm_source=google&utm_medium=cpc&utm_campaign=shop_google_performance_max_car_goods&gclid=CjwKCAjwvIWzBhAlEiwAHHWgvUw7z22sNzQYgdx6ORJkdTYXNLSwsYvaY52s8N399ZGVwCP5TbAt6xoCaD8QAvD_BwE"

	sum, err := extractPrices(url)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Сумма цен:", sum)
}

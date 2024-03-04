package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func Price(url string) string {
	// Run Chrome browser
	service, err := selenium.NewChromeDriverService("C:/chromedriver-win64/chromedriver.exe", 4444)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		// "--headless",  // comment out this line to see the browser
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}

	driver.Get(url)

	// Ждем несколько секунд для полной загрузки страницы (вы можете настроить под свои нужды)
	time.Sleep(5 * time.Second)

	// Находим элемент по классу
	elem, err := driver.FindElement(selenium.ByClassName, "item__price-left-side")
	if err != nil {
		panic(err)
	}

	// Получаем текст из элемента
	text, err := elem.Text()
	if err != nil {
		panic(err)
	}

	fmt.Println("Текст из элемента item__price-left-side:", text)
	return text
}

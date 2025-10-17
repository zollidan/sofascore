package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)


func GetUserDateChoice() (string, error) {
	menuItems := []string{
		"1. Вчера",
		"2. Сегодня",
		"3. Завтра",
	}

	fmt.Println("Выберите за какой период парсить данные:")
	for _, item := range menuItems {
		fmt.Println(item)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nВведите номер: ")
	scanner.Scan()

	choice, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "", fmt.Errorf("введите корректное число")
	}

	date, err := getDateByChoice(choice)
	if err != nil {
		return "", err
	}

	return date, nil
}

func getDateByChoice(choice int) (string, error) {
	now := time.Now()
	
	var date time.Time
	var label string

	switch choice {
	case 1:
		date = now.AddDate(0, 0, -1)
		label = "Вчера"
	case 2:
		date = now
		label = "Сегодня"
	case 3:
		date = now.AddDate(0, 0, 1)
		label = "Завтра"
	default:
		return "", fmt.Errorf("неверный выбор")
	}

	fmt.Printf("✓ Выбрано: %s (%s)\n\n", label, date.Format("02.01.2006"))
	return date.Format("2006-01-02"), nil
}

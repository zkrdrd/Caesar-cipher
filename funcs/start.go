package funcs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/***
Start - начало програмы, интеграционное меню
***/

func Start() {
	var (
		selecter int32
		shift    int32
		temp     string
	)
	get_text := bufio.NewScanner(os.Stdin)
	// обработка ввода выбора действий
	for {
		fmt.Print("Выбирите действие\n1 - зашифровать\n2 - расшифровать\n3 - найти смещение\n0 - выход\n:")
		fmt.Scanf("%s\n", &temp)
		if temp != "" {
			i, err := strconv.ParseInt(temp, 10, 32)
			if err == nil {
				selecter = int32(i)
				break
			} else {
				fmt.Println("Это должно быть целое число из доступных (1, 2, 3, 0).")
			}
		}
	}

	if selecter == 0 {
		os.Exit(0)
	} else if selecter != 3 {
		// обработка ввода смещения
		temp = ""
		for {
			fmt.Print("Введите смещение (Целое число больше 0): ")
			fmt.Scanf("%s\n", &temp)
			if temp != "" {
				i, err := strconv.ParseInt(temp, 10, 32)
				if err == nil {
					shift = int32(i)
					break
				} else {
					fmt.Println("Это должно быть целое число больше нуля.")
				}
			}
		}
	}
	// обработка ввода строки
	for {
		temp = ""
		fmt.Print("Введите строку: ")
		get_text.Scan()
		temp = get_text.Text()
		if temp > "" {
			break
		}
	}

	regular(selecter, shift, temp)
}

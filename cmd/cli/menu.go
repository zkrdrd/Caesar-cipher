package cli

import (
	"bufio"
	"caesarcipher/pkg/data"
	"fmt"
	"os"
	"strconv"
)

/***
ShowMenu - начало програмы, интеграционное меню
***/

func ShowMenu() {
	get_text := bufio.NewScanner(os.Stdin)
	// обработка ввода выбора действий
	for {
		fmt.Print("Выбирите действие\n1 - зашифровать\n2 - расшифровать\n3 - найти смещение\n0 - выход\n:")
		fmt.Scanf("%s\n", &data.Temp)
		if data.Temp != "" {
			i, err := strconv.ParseInt(data.Temp, 10, 32)
			if err == nil {
				data.Selecter = int32(i)
				break
			} else {
				fmt.Println("Это должно быть целое число из доступных (1, 2, 3, 0).")
			}
		}
	}

	if data.Selecter == 0 {
		os.Exit(0)
	} else if data.Selecter != 3 {
		// обработка ввода смещения
		data.Temp = ""
		for {
			fmt.Print("Введите смещение (Целое число больше 0): ")
			fmt.Scanf("%s\n", &data.Temp)
			if data.Temp != "" {
				i, err := strconv.ParseInt(data.Temp, 10, 32)
				if err == nil {
					data.Shift = int32(i)
					break
				} else {
					fmt.Println("Это должно быть целое число больше нуля.")
				}
			}
		}
	}
	// обработка ввода строки
	for {
		data.Temp = ""
		fmt.Print("Введите строку: ")
		get_text.Scan()
		data.Temp = get_text.Text()
		if data.Temp > "" {
			break
		}
	}

	regular()
}

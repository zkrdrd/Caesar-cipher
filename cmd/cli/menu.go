package cli

import (
	caeser "caesarcipher"
	"fmt"
	"os"
	"strconv"
)

/***
ShowMenu - начало програмы, интеграционное меню
***/

func ShowMenu() {
	// обработка ввода выбора действий
	for {
		fmt.Print("Выбирите действие\n1 - зашифровать\n2 - расшифровать\n3 - найти смещение\n0 - выход\n:")
		fmt.Scanf("%s\n", &caeser.Temp)
		if caeser.Temp != "" {
			i, err := strconv.ParseInt(caeser.Temp, 10, 32)
			if err == nil {
				caeser.Selecter = int32(i)
				break
			} else {
				fmt.Println("Это должно быть целое число из доступных (1, 2, 3, 0).")
			}
		}
	}

	if caeser.Selecter == 0 {
		os.Exit(0)
	} else if caeser.Selecter != 3 {
		// обработка ввода смещения
		caeser.Temp = ""
		for {
			fmt.Print("Введите смещение (Целое число больше 0): ")
			fmt.Scanf("%s\n", &caeser.Temp)
			if caeser.Temp != "" {
				i, err := strconv.ParseInt(caeser.Temp, 10, 32)
				if err == nil {
					caeser.Shift = int32(i)
					break
				} else {
					fmt.Println("Это должно быть целое число больше нуля.")
				}
			}
		}
	}
	// обработка ввода строки
	for {
		caeser.Temp = ""
		fmt.Print("Введите строку: ")
		caeser.GetText.Scan()
		caeser.Temp = caeser.GetText.Text()
		if caeser.Temp > "" {
			break
		}
	}

	regular()
}

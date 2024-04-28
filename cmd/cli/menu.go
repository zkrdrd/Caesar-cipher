package cli

import (
	"bufio"
	caeser "caesarcipher"
	"fmt"
	"os"
	"strconv"
)

var (
	// выбодействия зашифровка / расшифровка / поиск шага шифрования
	Selecter int32
	// переменная для интерактивного меню и передачи текста для шифрования / дешифрования
	Temp string
	// сканер
	GetText = bufio.NewScanner(os.Stdin)
)

// начало програмы, интеграционное меню
func ShowMenu() {
	// обработка ввода выбора действий
	for {
		fmt.Print("Выбирите действие\n1 - зашифровать\n2 - расшифровать\n3 - найти смещение\n0 - выход\n:")
		fmt.Scanf("%s\n", &Temp)
		if Temp != "" {
			i, err := strconv.ParseInt(Temp, 10, 32)
			if err == nil {
				Selecter = int32(i)
				break
			} else {
				fmt.Println("Это должно быть целое число из доступных (1, 2, 3, 0).")
			}
		}
	}

	if Selecter == 0 {
		os.Exit(0)
	} else if Selecter != 3 {
		// обработка ввода смещения
		Temp = ""
		for {
			fmt.Print("Введите смещение (Целое число больше 0): ")
			fmt.Scanf("%s\n", &Temp)
			if Temp != "" {
				i, err := strconv.ParseInt(Temp, 10, 32)
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
		Temp = ""
		fmt.Print("Введите строку: ")
		GetText.Scan()
		caeser.String = GetText.Text()
		if caeser.String > "" {
			break
		}
	}

	regular()
}

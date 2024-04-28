package cli

import (
	caeser "caesarcipher"
	"fmt"
	"os"
	"strconv"
)

// обработка введеных ранее данных
func regular() {
	if Selecter == 1 {
		d := caeser.Encoding(caeser.Shift, caeser.String)
		fmt.Println(d)
	} else if Selecter == 2 {
		d := caeser.Decoding(caeser.Shift, caeser.String)
		fmt.Println(d)
	} else {
		caeser.Finder(caeser.String)
	}
	now()
}

// интерактивное меню
func now() string {
	for {
		fmt.Print("\nВыбирите действие\n1 - еще раз\n0 - выход\n:")
		fmt.Scanf("%s\n", &Temp)
		i, err := strconv.ParseInt(Temp, 10, 32)
		if err == nil {
			if i == 1 {
				ShowMenu()
			}
			if i == 0 {
				os.Exit(0)
			}
		}
	}
}

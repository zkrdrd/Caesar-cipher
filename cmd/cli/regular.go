package cli

import (
	caeser "caesarcipher"
	"fmt"
	"os"
	"strconv"
)

// обработка введеных ранее данных
func regular() {
	varsd := caeser.Variables{}
	for _, chars := range caeser.Temp {
		if caeser.RuUpperLeft <= chars && chars <= caeser.RuUpperRight || caeser.RuLowerLeft <= chars && chars <= caeser.RuLowerRight {
			varsd = caeser.Variables{UpperLeft: caeser.RuUpperLeft,
				UpperRight:   caeser.RuUpperRight,
				LowerLeft:    caeser.RuLowerLeft,
				LowerRight:   caeser.RuLowerRight,
				AlphabetLen:  caeser.RuAlphabetLenght,
				EncDecShift:  caeser.Shift,
				EncDecString: caeser.Temp}
			break
		}

		if caeser.EnLowerLeft <= chars && chars <= caeser.EnLowerRight || caeser.EnUpperLeft <= chars && chars <= caeser.EnUpperRight {
			varsd = caeser.Variables{UpperLeft: caeser.EnUpperLeft,
				UpperRight:   caeser.EnUpperRight,
				LowerLeft:    caeser.EnLowerLeft,
				LowerRight:   caeser.EnLowerRight,
				AlphabetLen:  caeser.EnAlphabetLenght,
				EncDecShift:  caeser.Shift,
				EncDecString: caeser.Temp}
			break
		}
	}
	if caeser.Selecter == 1 {
		caeser.Encoding(varsd)
	} else if caeser.Selecter == 2 {
		caeser.Decoding(varsd)
	} else {
		caeser.Finder(varsd)
	}
	now()
}

// интерактивное меню
func now() string {
	for {
		fmt.Print("\nВыбирите действие\n1 - еще раз\n0 - выход\n:")
		fmt.Scanf("%s\n", &caeser.Temp)
		i, err := strconv.ParseInt(caeser.Temp, 10, 32)
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

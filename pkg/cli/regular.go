package cli

import (
	"caesarcipher/pkg/data"
	"fmt"
	"os"
	"strconv"
)

/***
regular - обработка введеных ранее данных
now 	- интерактивное меню
***/

func regular() {
	varsd := data.Variables{}
	for _, chars := range data.Temp {
		if data.RuUpperLeft <= chars && chars <= data.RuUpperRight || data.RuLowerLeft <= chars && chars <= data.RuLowerRight {
			varsd = data.Variables{UpperLeft: data.RuUpperLeft,
				UpperRight:   data.RuUpperRight,
				LowerLeft:    data.RuLowerLeft,
				LowerRight:   data.RuLowerRight,
				AlphabetLen:  data.RuAlphabetLenght,
				EncDecShift:  data.Shift,
				EncDecString: data.Temp}
			break
		}

		if data.EnLowerLeft <= chars && chars <= data.EnLowerRight || data.EnUpperLeft <= chars && chars <= data.EnUpperRight {
			varsd = data.Variables{UpperLeft: data.EnUpperLeft,
				UpperRight:   data.EnUpperRight,
				LowerLeft:    data.EnLowerLeft,
				LowerRight:   data.EnLowerRight,
				AlphabetLen:  data.EnAlphabetLenght,
				EncDecShift:  data.Shift,
				EncDecString: data.Temp}
			break
		}
	}
	if data.Selecter == 1 {
		Encoding(varsd)
	} else if data.Selecter == 2 {
		Decoding(varsd)
	} else {
		Finder(varsd)
	}
	now()
}

func now() string {
	for {
		fmt.Print("\nВыбирите действие\n1 - еще раз\n0 - выход\n:")
		fmt.Scanf("%s\n", &data.Temp)
		i, err := strconv.ParseInt(data.Temp, 10, 32)
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

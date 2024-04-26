package funcs

import (
	"caesarcipher/data"
	"fmt"
	"os"
	"strconv"
)

/***
regular - обработка введеных ранее данных
now 	- интерактивное меню
***/

func regular(selecter, shift int32, temp string) {
	varsd := data.Variables{}
	for _, chars := range temp {
		if data.RuUpperLeft <= chars && chars <= data.RuUpperRight || data.RuLowerLeft <= chars && chars <= data.RuLowerRight {
			varsd = data.Variables{UpperLeft: data.RuUpperLeft,
				UpperRight:   data.RuUpperRight,
				LowerLeft:    data.RuLowerLeft,
				LowerRight:   data.RuLowerRight,
				AlphabetLen:  data.RuAlphabetLenght,
				EncDecShift:  shift,
				EncDecString: temp}
			break
		}

		if data.EnLowerLeft <= chars && chars <= data.EnLowerRight || data.EnUpperLeft <= chars && chars <= data.EnUpperRight {
			varsd = data.Variables{UpperLeft: data.EnUpperLeft,
				UpperRight:   data.EnUpperRight,
				LowerLeft:    data.EnLowerLeft,
				LowerRight:   data.EnLowerRight,
				AlphabetLen:  data.EnAlphabetLenght,
				EncDecShift:  shift,
				EncDecString: temp}
			break
		}
	}
	if selecter == 1 {
		Encoding(varsd)
	} else if selecter == 2 {
		Decoding(varsd)
	} else {
		Finder(varsd)
	}
	now()
}

func now() string {
	var temp string
	for {
		fmt.Print("\nВыбирите действие\n1 - еще раз\n0 - выход\n:")
		fmt.Scanf("%s\n", &temp)
		i, err := strconv.ParseInt(temp, 10, 32)
		if err == nil {
			if i == 1 {
				Start()
			}
			if i == 0 {
				os.Exit(0)
			}
		}
	}
}

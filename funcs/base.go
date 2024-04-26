package funcs

import (
	"caesarcipher/data"
	"fmt"
)

/***
Encoding - коидрование строки
Decoding - декодирование строки
Finder	 - Поиск шага строки
***/

func Encoding(v data.Variables) {
	for _, chars := range v.EncDecString {
		// upper letter
		if v.UpperLeft <= chars && chars <= v.UpperRight {
			chars = int32(chars) + v.EncDecShift
			if chars > v.UpperRight {
				chars = chars - v.AlphabetLen
			}
		} else if v.LowerLeft <= chars && chars <= v.LowerRight { // lower letter
			chars = int32(chars) + v.EncDecShift
			if chars > v.LowerRight {
				chars -= v.AlphabetLen
			}
		}
		fmt.Printf("%c", chars)
	}
}

func Decoding(v data.Variables) {
	for _, chars := range v.EncDecString {
		// upper letter
		if v.UpperLeft <= chars && chars <= v.UpperRight {
			chars = int32(chars) - v.EncDecShift
			if chars < v.UpperLeft {
				chars += v.AlphabetLen
			}
		} else if v.LowerLeft <= chars && chars <= v.LowerRight { // lower letter
			chars = int32(chars) - v.EncDecShift
			if chars < v.LowerLeft {
				chars += v.AlphabetLen
			}
		}
		fmt.Printf("%c", chars)
	}
}

func Finder(v data.Variables) {
	var iteration int32
	for iteration = 1; iteration < v.AlphabetLen; iteration++ {
		fmt.Printf("%d - ", iteration)
		v.EncDecShift = iteration
		Decoding(v)
		fmt.Println()
	}
}

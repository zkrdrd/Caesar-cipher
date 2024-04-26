package caeser

import (
	"caesarcipher/pkg/data"
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
	for data.Iteration = 1; data.Iteration < v.AlphabetLen; data.Iteration++ {
		fmt.Printf("%d - ", data.Iteration)
		v.EncDecShift = data.Iteration
		Decoding(v)
		fmt.Println()
	}
}

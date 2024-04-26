package caeser

import (
	"fmt"
)

// Коидрование строки
func Encoding(v Variables) {
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

// Декодирование строки
func Decoding(v Variables) {
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

// Поиск шага строки
func Finder(v Variables) {
	for Iteration = 1; Iteration < v.AlphabetLen; Iteration++ {
		fmt.Printf("%d - ", Iteration)
		v.EncDecShift = Iteration
		Decoding(v)
		fmt.Println()
	}
}

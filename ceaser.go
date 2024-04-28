package caesarcipher

import (
	"fmt"
)

const (
	// левая граница русских символов алфавита в верхнем ресистре в таблице ascii
	RuUpperLeft = 1040
	// правая граница русских символов алфавита в верхнем ресистре в таблице ascii
	RuUpperRight = 1071
	// левая граница русских символов алфавита в нижнем регистре ресистре в таблице ascii
	RuLowerLeft = 1072
	// правая граница русских символов алфавита в нижнем ресистре в таблице ascii
	RuLowerRight = 1103
	// длина русского алфавита 32 символа из за того что буква 'ё' не учитывается
	RuAlphabetLenght = 32
	// левая граница английских символов алфавита в верхнем ресистре в таблице ascii
	EnUpperLeft = 65
	// правая граница английских символов алфавита в верхнем ресистре в таблице ascii
	EnUpperRight = 90
	// левая граница английских символов алфавита в нижнем ресистре в таблице ascii
	EnLowerLeft = 97
	// правая граница английских символов алфавита в нижнем ресистре в таблице ascii
	EnLowerRight = 122
	// длина аглийского алфавита
	EnAlphabetLenght = 26
)

var (
	// строка шифрования / дешифрования
	String string
	// шаг цикла
	Iteration int32
	// шаг шифрования
	Shift int32
	// результат
)

// Коидрование строки
func Encoding(Shift int32, String string) string {
	var Result string
	for _, chars := range String {
		// en upper letter
		if EnUpperLeft <= chars && chars <= EnUpperRight {
			chars = chars + Shift
			if chars > EnUpperRight {
				chars -= EnAlphabetLenght
			}
		} else if EnUpperLeft <= chars && chars <= EnLowerRight { // en lower letter
			chars = chars + Shift
			if chars > EnLowerRight {
				chars -= EnAlphabetLenght
			}
		} else if RuUpperLeft <= chars && chars <= RuUpperRight { // ru upper letter
			chars = chars + Shift
			if chars > RuUpperRight {
				chars -= RuAlphabetLenght
			}
		} else if RuLowerLeft <= chars && chars <= RuLowerRight { // ru lower letter
			chars = chars + Shift
			if chars > RuLowerRight {
				chars -= RuAlphabetLenght
			}
		}
		Result += string(chars)
	}
	return Result
}

// Декодирование строки
func Decoding(Shift int32, String string) string {
	var Result string
	for _, chars := range String {
		//  en upper letter
		if EnUpperLeft <= chars && chars <= EnUpperRight {
			chars = chars - Shift
			if chars < EnUpperLeft {
				chars += EnAlphabetLenght
			}
		} else if EnLowerLeft <= chars && chars <= EnLowerRight { // en lower letter
			chars = chars - Shift
			if chars < EnLowerLeft {
				chars += EnAlphabetLenght
			}
		} else if RuUpperLeft <= chars && chars <= RuUpperRight { // ru upper letter
			chars = chars - Shift
			if chars < RuUpperLeft {
				chars += RuAlphabetLenght
			}
		} else if RuLowerLeft <= chars && chars <= RuLowerRight { // ru lower letter
			chars = chars - Shift
			if chars < RuLowerLeft {
				chars += RuAlphabetLenght
			}
		}
		Result += string(chars)
	}
	return Result
}

// Поиск шага строки
func Finder(String string) {
	for Iteration = 1; Iteration < RuAlphabetLenght; Iteration++ {
		c := Decoding(Iteration, String)
		fmt.Printf("%d - %s\n", Iteration, c)
	}
}

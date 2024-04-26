package caeser

import (
	"bufio"
	"os"
)

/***
ascii table 	 - https://www.w3schools.com/charsets/ref_utf_cyrillic.asp
-----------------------------------------------------------------------------------------------------
RuUpperLeft 	 = левая граница русских символов алфавита в верхнем ресистре в таблице ascii
RuUpperRight 	 = правая граница русских символов алфавита в верхнем ресистре в таблице ascii
RuLowerLeft 	 = левая граница русских символов алфавита в нижнем регистре ресистре в таблице ascii
RuLowerRight 	 = правая граница русских символов алфавита в нижнем ресистре в таблице ascii
RuAlphabetLenght = длина русского алфавита 32 символа из за того что буква 'ё' не учитывается
-----------------------------------------------------------------------------------------------------
EnUpperLeft 	 = левая граница английских символов алфавита в верхнем ресистре в таблице ascii
EnUpperRight	 = правая граница английских символов алфавита в верхнем ресистре в таблице ascii
EnLowerLeft 	 = левая граница английских символов алфавита в нижнем ресистре в таблице ascii
EnLowerRight 	 = правая граница английских символов алфавита в нижнем ресистре в таблице ascii
EnAlphabetLenght = длина аглийского алфавита
***/

const (
	RuUpperLeft      = 1040
	RuUpperRight     = 1071
	RuLowerLeft      = 1072
	RuLowerRight     = 1103
	RuAlphabetLenght = 32
	EnUpperLeft      = 65
	EnUpperRight     = 90
	EnLowerLeft      = 97
	EnLowerRight     = 122
	EnAlphabetLenght = 26
)

/***
Iteration 	= итерация цикла
Selecter 	= выбодействия зашифровка / расшифровка / поиск шага шифрования
Shift	 	= шаг шифрования
Temp 		= переменная для интерактивного меню и передачи текста для шифрования / дешифрования
***/

var (
	Iteration int32
	Selecter  int32
	Shift     int32
	Temp      string
	GetText   = bufio.NewScanner(os.Stdin)
)

/***
UpperLeft 	 = левая граница символов алфавита в верхнем регистре передается из констант
UpperRight 	 = правая граница символов алфавита в верхнем регистре передается из констант
LowerLeft	 = левая граница символов алфавита в нижнем регистре передается из констант
LowerRight 	 = правая граница символов алфавита в нижнем регистре передается из констант
AlphabetLen  = длина алфамита передается из констант
EncDecShift	 = смещение символов
EncDecString = строка для шифрания / дешифрования
***/

type Variables struct {
	UpperLeft    int32
	UpperRight   int32
	LowerLeft    int32
	LowerRight   int32
	AlphabetLen  int32
	EncDecShift  int32
	EncDecString string
}

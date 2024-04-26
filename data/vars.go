package data

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

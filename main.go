package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/***
ascii table
https://www.w3schools.com/charsets/ref_utf_cyrillic.asp
***/

type variables struct {
	enc_dec        int32
	upper_left     int32
	upper_right    int32
	lower_left     int32
	lower_right    int32
	alphabet_len   int32
	enc_dec_shift  int32
	enc_dec_string string
}

func now() {
	var temp string
	for {
		fmt.Print("\nВыбирите действие\n1 - еще раз\n0 - выход\n:")
		fmt.Scanf("%s\n", &temp)
		i, err := strconv.ParseInt(temp, 10, 32)
		if err == nil {
			if i == 1 {
				main()
			}
			if i == 0 {
				os.Exit(0)
			}
		}
	}
}

func enc_dec_func(v variables) {
	for _, chars := range v.enc_dec_string {
		// upper letter
		if v.upper_left <= chars && chars <= v.upper_right {
			if v.enc_dec == 1 {
				chars = int32(chars) + v.enc_dec_shift
				if chars > v.upper_right {
					fmt.Printf("%c", chars-v.alphabet_len)
					continue
				}
				fmt.Printf("%c", chars)
			} else {
				chars = int32(chars) - v.enc_dec_shift
				if chars < v.upper_left {
					fmt.Printf("%c", chars+v.alphabet_len)
					continue
				}
				fmt.Printf("%c", chars)
			}

		} else if v.lower_left <= chars && chars <= v.lower_right { // lower letter
			if v.enc_dec == 1 {
				chars = int32(chars) + v.enc_dec_shift
				if chars > v.lower_right {
					fmt.Printf("%c", chars-v.alphabet_len)
					continue
				}
				fmt.Printf("%c", chars)
			} else {
				chars = int32(chars) - v.enc_dec_shift
				if chars < v.upper_left {
					fmt.Printf("%c", chars+v.alphabet_len)
					continue
				}
				fmt.Printf("%c", chars)
			}

		} else { //chars >= 32 && chars <= 47 {
			fmt.Printf("%c", chars)
		}
	}
}

func finder_func(v variables) {
	var iteration int32
	for iteration = 1; iteration < v.alphabet_len; iteration++ {
		fmt.Printf("%d - ", iteration)
		v.enc_dec_shift = iteration
		enc_dec_func(v)
		fmt.Println()
	}
}

func main() {
	var selecter int32
	var enc_dec_shift int32
	var enc_dec_string string
	var temp string
	get_text := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Выбирите действие\n1 - зашифровать\n2 - расшифровать\n3 - найти смещение\n0 - выход\n:")
		fmt.Scanf("%s\n", &temp)
		if temp != "" {
			i, err := strconv.ParseInt(temp, 10, 32)
			if err == nil {
				selecter = int32(i)
				break
			}
		}
	}

	if selecter == 0 {
		os.Exit(0)
	}

	if selecter == 1 || selecter == 2 {
		temp = ""
		for {
			fmt.Print("Введите смещение (Целое число > 0): ")
			fmt.Scanf("%s\n", &temp)
			if temp != "" {
				i, err := strconv.ParseInt(temp, 10, 32)
				if err == nil {
					enc_dec_shift = int32(i)
					break
				}
			}
		}
	}

	for {
		fmt.Print("Введите строку: ")
		get_text.Scan()
		enc_dec_string = get_text.Text()
		if enc_dec_string > "" {
			break
		}
	}
	var vars variables
	for _, chars := range enc_dec_string {
		if 1040 <= chars && chars <= 1071 || 1072 <= chars && chars <= 1103 {
			vars = variables{upper_left: 1040,
				upper_right:    1071,
				lower_left:     1072,
				lower_right:    1103,
				alphabet_len:   32,
				enc_dec_string: enc_dec_string,
				enc_dec_shift:  enc_dec_shift,
				enc_dec:        selecter}
			break
		}

		if 97 <= chars && chars <= 122 || 65 <= chars && chars <= 90 {
			vars = variables{upper_left: 65,
				upper_right:    90,
				lower_left:     97,
				lower_right:    122,
				alphabet_len:   26,
				enc_dec_string: enc_dec_string,
				enc_dec_shift:  enc_dec_shift,
				enc_dec:        selecter}
			break
		}
	}
	if selecter == 3 {
		finder_func(vars)
	} else {
		enc_dec_func(vars)
	}
	now()
}

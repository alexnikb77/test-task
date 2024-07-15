package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a []string
var k, p, l, f int
var z float64
var x, y, q string

func delSpaces(str string) string {
	result := make([]rune, 0, len(str))
	for _, char := range str {
		if char != ' ' {
			result = append(result, char)
		}
	}
	return string(result)
}

func findOper(st string) int {
	var i int
	for _, char := range st {
		switch char {
		case '+':
			{
				i++
				k = 1
			}
		case '-':
			{
				i++
				k = 2
			}

		case '*':
			{
				i++
				k = 3
			}

		case '/':
			{
				i++
				k = 4
			}

		}
	}
	//fmt.Println(i)
	return i
}

func mathcalc(one, two string) float64 {

	var res, one1, two1 float64
	var r, o int
	switch one {
	case "I":
		one1 = 1
	case "II":
		one1 = 2
	case "III":
		one1 = 3
	case "IV":
		one1 = 4
	case "V":
		one1 = 5
	case "VI":
		one1 = 6
	case "VII":
		one1 = 7
	case "VIII":
		one1 = 8
	case "IX":
		one1 = 9
	case "X":
		one1 = 10
	default:
		f = 1
	}
	switch two {
	case "I":
		two1 = 1
	case "II":
		two1 = 2
	case "III":
		two1 = 3
	case "IV":
		two1 = 4
	case "V":
		two1 = 5
	case "VI":
		two1 = 6
	case "VII":
		two1 = 7
	case "VIII":
		two1 = 8
	case "IX":
		two1 = 9
	case "X":
		two1 = 10
	default:
		f = 1
	}

	if f == 0 {
		switch k {
		case 1:
			res = float64(one1 + two1)
		case 2:
			res = float64(one1 - two1)
		case 3:
			res = float64(one1 * two1)
		case 4:
			res = float64(one1 / two1)
		}
		if res < 1 {
			//fmt.Println("По тз пример с римскими цифрами не может быть отридцательным ")
			//os.Exit(3)
			panic("По тз пример с римскими цифрами не может быть отридцательным и равным нулю")
		}
	} else {
		r, _ = strconv.Atoi(one)
		o, _ = strconv.Atoi(two)
		one1 = float64(r)
		two1 = float64(o)

		if one1 < 1 || two1 < 1 {
			//fmt.Println("пример не соответствует шаблону")
			//os.Exit(2)
			panic("пример не соответствует шаблону")
		}
		if one1 > 10 || two1 > 10 {
			//fmt.Println("пример не соответствует шаблону")
			//os.Exit(2)
			panic("пример не соответствует шаблону")
		}
		switch k {
		case 1:
			res = float64(one1 + two1)
		case 2:
			res = float64(one1 - two1)
		case 3:
			res = float64(one1 * two1)
		case 4:
			res = float64(one1 / two1)
		}
	}

	return res
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите пример")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = delSpaces(text)
		p = findOper(text)
		if p != 1 {
			//fmt.Println("пример не соответствует шаблону")
			//os.Exit(1)
			panic("пример не соответствует шаблону")
		}

		switch k {
		case 1:
			a = strings.Split(text, "+")
		case 2:
			a = strings.Split(text, "-")
		case 3:
			a = strings.Split(text, "*")
		case 4:
			a = strings.Split(text, "/")

		}

		z = mathcalc(a[0], a[1])
		l = int(z)
		if f == 0 {
			switch l {
			case 1:
				q = "I"
			case 2:
				q = "II"
			case 3:
				q = "III"
			case 4:
				q = "IV"
			case 5:
				q = "V"
			case 6:
				q = "VI"
			case 7:
				q = "VII"
			case 8:
				q = "VIII"
			case 9:
				q = "IX"
			case 10:
				q = "X"
			case 11:
				q = "XI"
			case 12:
				q = "XII"
			case 13:
				q = "XIII"
			case 14:
				q = "XIV"
			case 15:
				q = "XV"
			case 16:
				q = "XVI"
			case 17:
				q = "XVII"
			case 18:
				q = "XVIII"
			case 19:
				q = "IX"
			case 20:
				q = "XX"
			case 21:
				q = "XXI"
			case 22:
				q = "XXII"
			case 23:
				q = "XXIII"
			case 24:
				q = "XXIV"
			case 25:
				q = "XXV"
			case 26:
				q = "XXVI"
			case 27:
				q = "XXVII"
			case 28:
				q = "XXVIII"
			case 29:
				q = "XXIX"
			case 30:
				q = "XXX"
			case 31:
				q = "XXXI"
			case 32:
				q = "XXXII"
			case 33:
				q = "XXXIII"
			case 34:
				q = "XXXIV"
			case 35:
				q = "XXXV"
			case 36:
				q = "XXXVI"
			case 37:
				q = "XXXVII"
			case 38:
				q = "XXXVIII"
			case 39:
				q = "XXXIX"
			case 40:
				q = "XL"
			case 41:
				q = "XLI"
			case 42:
				q = "XLII"
			case 43:
				q = "XLIII"
			case 44:
				q = "XLIV"
			case 45:
				q = "XLV"
			case 46:
				q = "XLVI"
			case 47:
				q = "XLVII"
			case 48:
				q = "XLVIII"
			case 49:
				q = "XLIX"
			case 50:
				q = "L"
			case 51:
				q = "LI"
			case 52:
				q = "LII"
			case 53:
				q = "LIII"
			case 54:
				q = "LIV"
			case 55:
				q = "LV"
			case 56:
				q = "LVI"
			case 57:
				q = "LVII"
			case 58:
				q = "LVIII"
			case 59:
				q = "LVIX"
			case 60:
				q = "LX"
			case 61:
				q = "LXI"
			case 62:
				q = "LXII"
			case 63:
				q = "LXIII"
			case 64:
				q = "LXIV"
			case 65:
				q = "LXV"
			case 66:
				q = "LXVI"
			case 67:
				q = "LXVII"
			case 68:
				q = "LXVIII"
			case 69:
				q = "LXIX"
			case 70:
				q = "LXX"
			case 71:
				q = "LXXI"
			case 72:
				q = "LXXII"
			case 73:
				q = "LXXIII"
			case 74:
				q = "LXXIV"
			case 75:
				q = "LXXV"
			case 76:
				q = "LXXVI"
			case 77:
				q = "LXXVII"
			case 78:
				q = "LXXVIII"
			case 79:
				q = "LXXIX"
			case 80:
				q = "LXXX"
			case 81:
				q = "LXXXI"
			case 82:
				q = "LXXXII"
			case 83:
				q = "LXXXIII"
			case 84:
				q = "LXXXIV"
			case 85:
				q = "LXXXV"
			case 86:
				q = "LXXXVI"
			case 87:
				q = "LXXXVII"
			case 88:
				q = "LXXXVIII"
			case 89:
				q = "LXXIX"
			case 90:
				q = "XC"
			case 91:
				q = "XCI"
			case 92:
				q = "XCII"
			case 93:
				q = "XCIII"
			case 94:
				q = "XCIV"
			case 95:
				q = "XCV"
			case 96:
				q = "XCVI"
			case 97:
				q = "XCVII"
			case 98:
				q = "XCVIII"
			case 99:
				q = "XCVIX"
			case 100:
				q = "C"
			}
			fmt.Println(q)
		} else {
			fmt.Println(l)
		}

	}

}

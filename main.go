package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a []string
var k, p int
var z float64
var x, y string

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
	var f, x, y int
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
		if res < 0 {
			//fmt.Println("По тз пример с римскими цифрами не может быть отридцательным ")
			//os.Exit(3)
			panic("По тз пример с римскими цифрами не может быть отридцательным")
		}
	} else {
		x, _ = strconv.Atoi(one)
		y, _ = strconv.Atoi(two)
		one1 = float64(x)
		two1 = float64(y)

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
		fmt.Println(z)

	}

}

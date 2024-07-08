package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a []string
var k, x, y, p int
var z float64

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
		default:

		}
	}
	//fmt.Println(i)
	return i
}

func mathcalc(one, two int) float64 {

	var res, one1, two1 float64
	one1 = float64(one)
	two1 = float64(two)
	if one < 1 || two < 1 {
		fmt.Println("пример не соответствует шаблону")
		os.Exit(2)
	}
	if one > 10 || two > 10 {
		fmt.Println("пример не соответствует шаблону")
		os.Exit(2)
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
			fmt.Println("пример не соответствует шаблону")
			os.Exit(1)
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
		x, _ = strconv.Atoi(a[0])
		y, _ = strconv.Atoi(a[1])

		z = mathcalc(x, y)
		//fmt.Println(y)
		//fmt.Println(x)

		fmt.Println(z)

	}

}

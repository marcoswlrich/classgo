package funcoes

import "fmt"

func Calculos() {
	var num3 int = 30
	var num4 int = 240

	calculo := func(num1 int, num2 int) int {
		return num1 + num2 + num3 + num4
	}

	fmt.Println(calculo(10, 25))

	calculo = func(num1 int, num2 int) int {
		return num1 * num2 * num3
	}
	fmt.Println(calculo(10, 25))
}

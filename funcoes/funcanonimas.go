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

func sum(a, b int) int {
	return a + b
}

func multi(a, b int) (int, bool) {
	if a*b >= 100 {
		return a * b, true
	}
	return a * b, false
}

// variadicas
func sum2(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

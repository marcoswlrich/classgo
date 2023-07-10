package middleware

import "fmt"

func somar(a, b int) int {
	return a + b
}

func subtrair(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MeuMiddleware() {
	fmt.Println("Inicio")

	result := operaMidd(somar)(2, 2)
	result = operaMidd(subtrair)(14, 7)
	fmt.Println(result)
	result = operaMidd(multiplicar)(5, 5)
	fmt.Println(result)
}

func operaMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacao")
		return f(a, b)
	}
}

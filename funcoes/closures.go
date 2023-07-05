package funcoes

import "fmt"

func tabela(valor int) func() int {
	num := valor
	sequencia := 0
	return func() int {
		sequencia++
		return num * sequencia
	}
}

func ChamarClosure() {
	tabela2 := 2
	MinhaTabela := tabela(tabela2)
	for i := 1; i <= 10; i++ {
		fmt.Println(MinhaTabela())
	}
}

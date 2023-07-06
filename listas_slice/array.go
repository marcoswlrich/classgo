package listas_slice

import "fmt"

var (
	tabela [10]int = [10]int{10, 0, 59} // add
	matriz [10][20]int
)

func MostraArrays() {
	tabela[7] = 33
	tabela[2] = 54
	fmt.Println(tabela)

	for i := 0; i < len(tabela); i++ {
		fmt.Println(tabela[i])
	}

	tabela2 := [10]string{"Marcos", "Pedro"}

	fmt.Println(tabela2)

	matriz[4][14] = 15
	println(matriz)
}

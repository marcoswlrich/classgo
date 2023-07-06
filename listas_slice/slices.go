package listas_slice

import "fmt"

var (
	tabelaS []int   = []int{2, 5, 4}
	arrays  [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8}
)

func MostraSlice() {
	fmt.Println(tabelaS)

	position := arrays[3:]   // Slice criado a partir de um vetor, a partir da posicao position3
	position2 := arrays[:5]  // Slice criado a partir de um vetor, a partir da posicao 0 ate 5
	position3 := arrays[6:7] // Slice criado a partir de um vetor, desde a posicap 6 ate 7

	fmt.Println(position)
	fmt.Println(position2)
	fmt.Println(position3)
}

func Capacidade() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largura %d, Capacidade %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largura %d, Capacidade %d", len(nums), cap(nums))
}

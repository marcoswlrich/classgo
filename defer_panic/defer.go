package defer_panic

import (
	"fmt"
	"log"
)

func VerDefer() {
	fmt.Println("Primeira mensagem")
	defer fmt.Println("Ultima mensagem")

	fmt.Println("Segunda mensagem")
}

func ExPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocorreu um erro e gerou um Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 2 {
		panic("Encontrou o valor 2")
	}
}

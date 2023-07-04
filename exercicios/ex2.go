package exercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TabeladeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	var num int
	var err error
	var texto string

	for {
		fmt.Println("Adicione um numero: ")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", num, i, i*num)
	}

	return texto
}

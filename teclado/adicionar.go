package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	num1    int
	num2    int
	legenda string
	err     error
)

func AddNum() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o numero 1 : ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("O dado adicionado e incorreto " + err.Error())
		}
	}

	fmt.Println("Digite o numero 2 : ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("O dado adicionado e incorreto " + err.Error())
		}
	}

	fmt.Println("Digite a legenda : ")
	if scanner.Scan() {
		legenda = scanner.Text()
	}

	fmt.Println(legenda, num1*num2)
}

package variaveis

import (
	"fmt"
	"strconv"
	"time"
)

var (
	Nombre string
	Estado bool
	Saldo  float32
	Timer  time.Time
)

func RestoVariaveis() {
	Nombre = "Pedro"
	Estado = true
	Saldo = 1527.88
	Timer = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Saldo)
	fmt.Println(Timer)
}

func ConvertTexto(num int) (bool, string) {
	texto := strconv.Itoa(num)
	return true, texto
}

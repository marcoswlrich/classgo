package main

import (
	"fmt"

	"github.com/marcoswlrich/classgo/variaveis"
)

func main() {
	estado, texto := variaveis.ConvertTexto(1999)
	fmt.Println(estado)
	fmt.Println(texto)
}

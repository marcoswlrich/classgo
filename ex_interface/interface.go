package ex_interface

import (
	"fmt"

	"github.com/marcoswlrich/classgo/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Sou um/a %s, e estou respirando \n", hu.Sexo())
}

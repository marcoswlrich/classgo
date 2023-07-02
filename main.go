package main

import (
	"fmt"
	"strconv"

	"github.com/marcoswlrich/classgo/exercicios"
	"github.com/marcoswlrich/classgo/teclado"
)

func main() {
	// estado, texto := variaveis.ConvertTexto(1999)
	// fmt.Println(estado)
	// fmt.Println(texto)
	// if os := runtime.GOOS; os == "linux" || os == "OS X." {
	// 	fmt.Println("Linux")
	// } else {
	// 	fmt.Println("Windows")
	// }
	//
	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Linux")
	// case "darwin":
	// 	fmt.Println("Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	//
	// }
	texto := "42.55"

	b, _ := strconv.ParseFloat(texto, 64)
	fmt.Printf("%T \n", b)
	fmt.Println(b)

	testEx, returnEx := exercicios.ConvNum("101")
	fmt.Println(testEx)
	fmt.Println(returnEx)

	teclado.AddNum()
}

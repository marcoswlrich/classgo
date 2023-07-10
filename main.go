package main

import (
	"github.com/marcoswlrich/classgo/webserver"
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
	// texto := "42.55"
	//
	// b, _ := strconv.ParseFloat(texto, 64)
	// fmt.Printf("%T \n", b)
	// fmt.Println(b)
	//
	// testEx, returnEx := exercicios.ConvNum("101")
	// fmt.Println(testEx)
	// fmt.Println(returnEx)
	//
	// teclado.AddNum()

	// files.GravaTabela()
	// files.AddTabela()
	// files.LerArquivo()

	// canal1 := make(chan bool)
	// go goroutines.NomeLento("Marcos Wlrich", canal1)
	//
	// defer func() {
	// 	<-canal1
	// }()
	// fmt.Println("Estou aqui")
	// // var x string
	// // fmt.Scanln(&x)
	//
	webserver.WebServer()
}

package main

import (
	"fmt"
	"runtime"

	"github.com/marcoswlrich/classgo/variaveis"
)

func main() {
	estado, texto := variaveis.ConvertTexto(1999)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Linux")
	} else {
		fmt.Println("Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("Darwin")
	default:
		fmt.Printf("%s \n", os)

	}
}

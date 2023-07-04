package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marcoswlrich/classgo/exercicios"
)

var fileName string = "./files/txt/tabela.txt"

func GravaTabela() {
	var texto string = exercicios.TabeladeMultiplicar()
	arquivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo" + err.Error())
		return
	}
	fmt.Fprintln(arquivo, texto)
	arquivo.Close()
}

func AddTabela() {
	var texto string = exercicios.TabeladeMultiplicar()
	if !Append(fileName, texto) {
		fmt.Println("Erro ao concatenar o conteudo")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Erro durante o Append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Erro durante o Append" + err.Error())
		return false
	}

	arch.Close()

	return true
}

func LerArquivo() {
	arquivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Erro ao ler arquivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	arquivo.Close()
}

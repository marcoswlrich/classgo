package structs

import (
	"fmt"
)

type Endereco struct {
	Logradouro string
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	Rua Endereco
}

func cliente() {
	marcos := Cliente{
		Nome:  "Marcos",
		Idade: 30,
		Ativo: true,
	}
	marcos.Cidade = "Cuiaba"

	marcos.Rua.Logradouro = "Rua 03"

	fmt.Println(marcos.Nome)
}

// Metodos em Structs

type Carro struct {
	Nome  string
	Marca string
	Ano   int
	Ativo bool
}

func (c Carro) Desativar() {
	c.Ativo = false
	fmt.Println(c.Ativo)
	fmt.Printf("O carro %s foi desativado", c.Nome)
}

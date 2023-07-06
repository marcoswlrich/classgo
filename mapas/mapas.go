package mapas

import "fmt"

func MostraMapas() {
	paises := make(map[string]string)

	paises["Brasil"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Brasil"])

	campeonato := map[string]int{
		"Palmeiras":    45,
		"Real Madri":   39,
		"River":        39,
		"Boca Juniors": 26,
	}

	fmt.Println(campeonato)

	for equipe, pontos := range campeonato {
		fmt.Printf("Equipe %s, tem %d pontos no campeonato \n", equipe, pontos)
	}

	delete(campeonato, "Boca Juniors")
	fmt.Println(campeonato)

	pontos, existe := campeonato["Corinthians"]
	fmt.Printf("Os pontos capturados es %d, a equipe existe = %t \n", pontos, existe)
}

package exercicios

import "strconv"

func ConvNum(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "error" + err.Error()
	}
	if num > 100 {
		return num, "E maior que 100"
	} else {
		return num, "E menor que 100"
	}
}

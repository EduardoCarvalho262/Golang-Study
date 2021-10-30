package main

import (
	"fmt"
)

func main() {
	var par, impar, positivo, negativo int
	numeros := make([]int, 5)

	fmt.Scan(&numeros[0], &numeros[1], &numeros[2], &numeros[3], &numeros[4])

	for _, valor := range numeros {
		if valor % 2 < 0 || valor % 2 == 1 {
			impar++
		}
		if valor % 2 == 0 {
			par++
		}
		if valor > 0 {
			positivo++
		}
		if valor < 0 {
			negativo++
		}
	}

	fmt.Printf("%d valor(es) par(es)\n", par)
	fmt.Printf("%d valor(es) impar(es)\n", impar)
	fmt.Printf("%d valor(es) positivo(s)\n", positivo)
	fmt.Printf("%d valor(es) negativo(s)\n", negativo)
}

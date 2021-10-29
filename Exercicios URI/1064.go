package main

import (
	"fmt"
)

func main() {
	numeros := make([]float64, 6)
	var contador int
	var media float64
	var soma float64

	fmt.Scan(&numeros[0], &numeros[1], &numeros[2], &numeros[3], &numeros[4], &numeros[5])

	for _, valor := range numeros {
		if valor > 0 {
			contador++
			soma += valor
		}
	}

	fmt.Printf("%d valores positivos\n", contador)
	media = soma / float64(contador)
	fmt.Printf("%.1f\n", media)
}

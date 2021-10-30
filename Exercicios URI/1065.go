package main

import (
	"fmt"
	"math"
)

func main() {

	var contador int

	numeros := make([]float64, 5)

	fmt.Scan(&numeros[0], &numeros[1], &numeros[2], &numeros[3], &numeros[4])

	for _, valor := range numeros {
		if math.Mod(valor, 2) == 0 {
			contador++
		}
	}
}
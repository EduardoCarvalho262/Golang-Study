package main

import (
	"fmt"
)

func main() {

	var cod1, num1, cod2, num2 int
	var valor1, valor2 float64

	fmt.Scan(&cod1, &num1, &valor1)
	fmt.Scan(&cod2, &num2, &valor2)

	x := (float64(num1) * valor1) + (float64(num2) * valor2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n",x)

}

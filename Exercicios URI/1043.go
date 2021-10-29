package main

import (
	"fmt"
)

func main() {
	var salario, imposto float64

	fmt.Scan(&salario)

	if salario <= 2000.00 {
		fmt.Println("Isento")
	} else if salario <= 3000.00 {
		imposto = (salario - 2000.00) * 0.08
		fmt.Printf("R$ %.2f\n", imposto)
	} else if salario <= 4500.00 {
		imposto = 1000.00 * 0.08 + (salario - 3000.00) * 0.18
		fmt.Printf("R$ %.2f\n", imposto)
	} else {
		imposto =  1000.00 * 0.08 + 1500.00 * 0.18 + (salario - 4500.00) * 0.28
		fmt.Printf("R$ %.2f\n", imposto)
	}

}

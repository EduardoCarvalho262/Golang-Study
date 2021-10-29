package main

import (
	"fmt"
)

func main() {

	var salario float64
	var vendasTotais float64
	var nome string

	fmt.Scan(&nome)
	fmt.Scan(&salario)
	fmt.Scan(&vendasTotais)
	aumento := vendasTotais * 0.15

	total := salario + aumento

	fmt.Printf("TOTAL = R$ %.2f", total)
	fmt.Println()

}

package main

import (
	"fmt"
	"math"
)

func main() {
	var idade int // declarando variavel
	fmt.Println("Minha idade é", idade)
	idade = 29 // atribuindo valor
	fmt.Println("Minha idade é", idade)

	//tipo por inferencia
	var idade2 = 30
	fmt.Println("Minha verdadeira é", idade2)

	//Multipla declaracao de variaveis
	var altura, largura = 1, 62.0
	fmt.Println("Altura é", altura, "a largura é", largura)

	var (
		nome      = "Ramon"
		sobrenome = "Lean"
	)
	fmt.Println("My name is", nome, "meu sobrenome é", sobrenome)

	//Atalho de declaracao de variavel
	count := 10
	fmt.Println("Contador é", count)

	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)

}

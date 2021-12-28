package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	//Configuração do Log
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Eduardo")

	//Verificando se retorna erro
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

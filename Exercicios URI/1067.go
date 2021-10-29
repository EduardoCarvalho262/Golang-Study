package main

import (
	"fmt"
)

func main() {
	var numero int

	fmt.Scan(&numero)

	for i := 1; i <= numero; i++ {
		if(i % 2 == 1){
			fmt.Println(i)
		}
	}
}

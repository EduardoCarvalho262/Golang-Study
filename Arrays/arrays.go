package main

import (
	"fmt"
)

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func main() {
	a := [3]int{12, 78, 50} // short hand declaration to create array
	b := [...]int{12, 36, 48, 9}
	fmt.Println(a)
	fmt.Println(b)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	fmt.Println("Length is", len(b))

	sum := int(0)
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %d\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)

	for i, valor := range b {
		fmt.Printf("O valor de i é %d, o valor de valor é %d \n", i, valor)
	}


	//Slice
	
}

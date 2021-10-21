package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
}

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

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
	fmt.Printf("\nCapacidade do slice %d\n", cap(dslice))

	i := make([]int, 5, 6)
	fmt.Println(i)

	//Variatic
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)

}

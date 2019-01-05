package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Println(a, b)

	// type of b is *int
	fmt.Printf("%T\n", b)

	//use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//change val with pointer
	*b = 10
	fmt.Println(a)
}
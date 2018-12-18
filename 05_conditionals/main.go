package main

import "fmt"

func main() {
	x := 12
	y := 10
	if x <= y {
		fmt.Printf("%d is less than or equal to %d!\n",x,y)
	} else {
		fmt.Printf("%d is less than %d!\n",y,x)
	}

	// else if
	color := "yellow"
	if color == "red" {
		fmt.Println("color is red!")
	} else if color == "yellow" {
		fmt.Println("color is yellow!")
	} else {
		fmt.Println("color is NOT red or yellow!")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT red or blue")
	}

}
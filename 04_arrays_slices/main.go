package main

import "fmt"

func main() {
	// var fruitArr [2]string
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"

    fruitArr := [2]string{"Apple","Orange"}
    fruitSlice := []string{"Apple","Orange","Grape","Cherry"}

	fmt.Println(fruitArr)
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[0:1])
	fmt.Println(fruitArr[0])
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:4])
}
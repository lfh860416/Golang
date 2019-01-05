package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("SUM", sum)

	//range with map
	emails := map[string]string{"Alex":"alex@gmail.com", "Winnie":"winnie@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range  emails {
		fmt.Println("Name: " + k)
	}
}
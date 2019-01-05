package main

import "fmt"

func main() {
	//define map
	emails := make(map[string]string)
	//assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	//declare map and add kv
	emails2 := map[string]string{"Alex":"alex@gmail.com", "Winnie":"winnie@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Mike"])
	fmt.Println("emails2 is", emails2)

	//delete from map
	delete(emails2, "Alex")
	fmt.Println("emails2 is", emails2)
}
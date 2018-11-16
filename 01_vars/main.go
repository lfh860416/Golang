package main


import "fmt"

func main() {
	var name1 = "winne"
	var name2 string = "jerry"
	//shorthand
	name3, email := "alex", "alex@163.com"

	//int
	var age = 32
	var age2 int32 = 32

	var size float32 = 2.3
	//float64
	size2 := 1.2

	var isCute = false
	const isCool = true

	fmt.Println(name3, age, email, size, isCool)
	fmt.Printf("%T %T %T %T\n", name1, name2, age, age2)
	fmt.Printf("%T %T %T %T\n", size, size2, isCute, isCool)

}

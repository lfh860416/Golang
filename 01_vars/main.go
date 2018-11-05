package main
import "fmt"

func main() {
	var name = "winne"
	//shorthand
	name2 := "jerry"
	name3, email := "alex", "alex@163.com"

	//int
	var age = 32
	var age2 int32 = 32

	var size float32 = 2.3
	//float64
	size2 := 1.2

	const isCool = true

	fmt.Println(name, name2, age, name3, email)
	fmt.Printf("%T %T %T %T %T\n", size, size2, age, age2, isCool)

}

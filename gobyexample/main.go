package main

import "fmt"
import "math"
import "time"

const s string = "constant"

// https://gobyexample.com/if-else
func main() {
	fmt.Println("hello world")

    //values
	fmt.Println("xcy" + " baby")
	fmt.Println("1+2=",1+2)
	fmt.Println("7.0/3.0=",7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//variables
	var a = "initial"
	fmt.Println(a)
	var b,c int = 1,2
	fmt.Println(b,c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	f := "short"
	fmt.Println(f)

	//constants
	fmt.Println(s)
	const n = 500000000
	const m = 3e20 / n
	fmt.Println(m)
	// a numeric constant has no type until it's given one
	fmt.Println(int64(m))
	fmt.Println(math.Sin(n))

	//For
	for i := 1;i <= 3; i++ {
		fmt.Println(i)
	}
	j := 5
	for j <= 7 {
		fmt.Println(j)
		j = j+1
	}
	for {
		fmt.Println("loop")
		break
	}
	//For
	for k := 0;k <= 3; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}	

	//If/Else
	if num := 9; num < 0 {
		fmt.Println(num, "is negative!")
	} else if n < 10 {
		fmt.Println(num, "has one digit!")
	} else {
		fmt.Println(num, "has multiple digits!")
	}

	//switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday!")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Unknown type %T\n", t)
		}	
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	
}
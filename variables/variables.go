package main

import "fmt"

func main() {
	a := "initial"
	fmt.Println(a)

	b, c := 1, 2 // multiple variables in one line
	fmt.Println(b, c)

	d := true
	fmt.Println(d)

	var e int // automatically initializes variable to 0 or the nil value for that type
	fmt.Println(e)

	f := "apple" // type inferred
	fmt.Println(f)
}

package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") // concatenate strings
	fmt.Println("1 + 1", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// boolean logic
	fmt.Println(true && false) // and
	fmt.Println(true || false) // or
	fmt.Println(!true)         // negation
}

package main

import "fmt"

// global variables
const pi float64 = 3.1416

var x int = 21

// global scope
// {local scope}

func main() {
	// Local variables
	// main's local scope
	a := 10
	var b int = 20

	c := a + b
	fmt.Println(c)

	if a > 10 {
		fmt.Println("a is greater than 10")
	} else if a <= 10 {
		fmt.Println("a is less than or equal to 10")
	} else {
		fmt.Println("a is not a valid input")
	}
	fmt.Println(x, pi)

	add()
}

func add() {
	// add's local scope
	a := 11
	res := a + x
	fmt.Println(res)
}

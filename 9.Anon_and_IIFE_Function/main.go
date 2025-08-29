package main

import "fmt"

// init() function
// Special function in Go that is called automatically BEFORE main().
// You cannot call it yourself, Go runtime does it.
// Useful for initialization tasks like setting variables, configs, or DB connections.
func init() {
	fmt.Println("I am the init function")
}

// Example of a standard/named function (commented for reference).
// func add(a, b int) {
// 	fmt.Println(a + b)
// }

func main() {
	// add(5, 7) // would call a named function if it were uncommented.

	// -------------------
	// Anonymous Function + IIFE
	// -------------------
	// - Anonymous function: A function without a name.
	// - IIFE: Immediately Invoked Function Expression.
	//   Declared and called at the same time.
	//   This avoids storing it in a variable when you just want to run it once.

	func(a int, b int) {
		c := a + b
		fmt.Println("IIFE Result:", c)
	}(5, 7) // <- Call immediately with (5, 7)
}

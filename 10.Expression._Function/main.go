package main

import "fmt"

// ----------------------
// Function Declarations
// ----------------------

// Named function without parameters
// - Declared at package scope
// - Can be called from anywhere in this package
func sum() {
	fmt.Println(10 + 11)
}

// Named function with parameters
// - Takes two ints, prints their sum
func add(a, b int) {
	fmt.Println(a + b)
}

// ----------------------
// Special init() Function
// ----------------------
// - Runs automatically before main().
// - Every package can have multiple init() functions.
// - Common use cases: initializing configs, connections, seeding data.
func init() {
	fmt.Println("I am the init function and I'm called first")
}

// ----------------------
// main() Function
// ----------------------
// - Entry point of the Go program.
// - Only one main() is allowed per package "main".
func main() {
	// 1. Calling a named function (declared at package scope)
	sum()

	// 2. Regular variable declaration inside block scope
	s := "we are finished"

	// 3. Calling a named function with arguments
	add(2, 3)

	// 4. Function Expression (Anonymous Function assigned to variable)
	//    - Here `add` shadows the outer package-scoped `add`
	//    - Demonstrates closures / function literals
	add := func(x int, y int) {
		c := x + y
		fmt.Println(c)
	}

	// Calling the function expression
	add(4, 4)

	// 5. Anonymous function stored in a variable
	//    - MUST be defined before use (unlike package-level functions).
	//    - If you try to call mul() before this assignment, it fails.
	// mul(10, 15) // Error: mul is not defined yet
	mul := func(x int, y int) {
		c := x * y
		fmt.Println(c)
	}
	// Works since mul is already defined
	mul(10, 15)

	// 6. Printing local variable
	fmt.Println(s)
}

package main

import "fmt"

//////////////////////
// FIRST ORDER FUNCTIONS
//////////////////////

// A first-order function is a "regular" function.
// It neither takes a function as a parameter nor returns one.
func add(a int, b int) { // parameters => a, b
	c := a + b
	fmt.Println("Sum:", c)
}

func multiply(a int, b int) {
	c := a * b
	fmt.Println("Multiplied:", c)
}

// ----------------------
// Higher Order Functions (HOF)
// ----------------------

/*
A function is considered a Higher Order Function if:
1. It takes another function as a parameter, OR
2. It returns a function, OR
3. It does both.

This allows Go to support a **functional style of programming**
even though Go is not a pure functional language.
*/

// (a) Function as parameter
func processOperation(a int, b int, operation func(x int, y int)) {
	// here `operation` is any function that takes (int, int)
	operation(a, b)
}

// (b) Function as return value
func call() func(x int, y int) {
	// returning an existing function (add)
	return add
}

// ----------------------
// MAIN FUNCTION
// ----------------------

func main() {
	// First Order Function
	add(2, 5) // arguments => 2, 5

	// Higher Order Function → Function as Parameter
	processOperation(3, 7, add)
	processOperation(3, 7, multiply)

	// Higher Order Function → Function as Return
	sum := call() // assigns the returned function to sum
	sum(4, 7)     // calls add(4, 7) via sum

	// Anonymous Function (First-order, but unnamed)
	anon := func(a, b int) {
		fmt.Println("Anonymous Multiplication:", a*b)
	}
	anon(2, 6)

	// IIFE (Immediately Invoked Function Expression)
	func(x, y int) {
		fmt.Println("IIFE Subtraction:", x-y)
	}(10, 3) // executed immediately
}

//////////////////////
// NOTES & SUMMARY
//////////////////////

/*
 KEY TERMS

1. Parameter vs Argument
   - Parameter → the variable names defined in the function signature (e.g., a, b)
   - Argument → the actual values passed during function call (e.g., 2, 5)

2. First Order Function
   - Standard named function (e.g., add)
   - Anonymous function (e.g., `func(a, b int) {}`)
   - IIFE (Immediately Invoked Function Expression)
   - Function expression (assign function to a variable)

3. Higher Order Function or First Class Function
   - Function that accepts another function as input (processOperation)
   - Function that returns another function (call)
   - Function that does both (can be built by combining patterns)
4. Callback function
	- When a higher order function takes another function as a parameter or the function we pass to a higher order function function as an argument
5. First Class Citizen
	- Any data assigned to a variable is called a first class citizen

Why Important?
- First-order functions = basic building blocks
- Higher-order functions = enable abstraction, reusable logic,
  and functional programming patterns (like map, filter, reduce in other languages)
*/

package main

import "fmt"

// 1. Standard or Named Function
// A normal function with a name that can be called multiple times.
func add(x int, y int) {
	fmt.Println("Standard Function:", x+y)
}

// 2. Anonymous Function
// A function without a name. It must be assigned or used immediately.
var anon = func(msg string) {
	fmt.Println("Anonymous Function:", msg)
}

// 3. Function Expression (Assign function to a variable)
// Functions in Go are first-class citizens, so they can be assigned to variables.
var multiply = func(x, y int) int {
	return x * y
}

// 4. Higher-Order Function (First-class function)
// A function that takes another function as an argument OR returns a function.
func operate(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

// 5. Callback Function
// Using a function passed as an argument (common in async programming).
func process(name string, callback func(string)) {
	callback(name)
}

// 6. Variadic Function
// Accepts variable number of arguments.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 7. init() Function
// Runs automatically before main().
// Cannot be called manually.
func init() {
	fmt.Println("Init Function: runs before main()")
}

// 8. Closure
// A closure "closes over" variables from its surrounding scope.
func closureExample() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 9. Defer Function
// Deferred functions are executed AFTER surrounding function finishes (LIFO).
func deferExample() {
	defer fmt.Println("Defer 1: runs last")
	defer fmt.Println("Defer 2: runs second")
	fmt.Println("Inside deferExample: runs first")
}

// 10. Receiver Function (Method on a type)
type Person struct {
	name string
}

// Method with receiver (like methods in OOP)
func (p Person) greet() {
	fmt.Println("Receiver Function:", "Hello, my name is", p.name)
}

// 11. IIFE (Immediately Invoked Function Expression)
// Declared and executed immediately.
var iife = func(msg string) string {
	return "IIFE: " + msg
}("Runs immediately!")

// ---------------- MAIN ----------------
func main() {
	// 1. Named Function
	add(5, 7)

	// 2. Anonymous Function (via variable)
	anon("Hello Go!")

	// 3. Function Expression
	fmt.Println("Function Expression:", multiply(3, 4))

	// 4. Higher Order Function
	result := operate(10, 20, multiply)
	fmt.Println("Higher Order Function:", result)

	// 5. Callback Function
	process("Sadik", func(n string) {
		fmt.Println("Callback Function: Hello,", n)
	})

	// 6. Variadic Function
	fmt.Println("Variadic Function:", sum(1, 2, 3, 4, 5))

	// 8. Closure Example
	closureCounter := closureExample()
	fmt.Println("Closure:", closureCounter()) // 1
	fmt.Println("Closure:", closureCounter()) // 2

	// 9. Defer Example
	deferExample()

	// 10. Receiver Function
	p := Person{name: "Sami"}
	p.greet()
}

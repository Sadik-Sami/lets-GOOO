package main

import "fmt"

// Global variable
var a = 10

// The init() function is special in Go.
// - It runs **before main()** automatically.
// - You cannot call it manually.
// - Each package can have multiple init() functions (even in different files).
func init() {
	fmt.Println("I am the init function, I can't be called. I am automatically called")
	fmt.Println("Value of a inside init():", a) // 10 (global value)

	// You can modify global variables inside init()
	a = 20
}

// The main() function is the entry point of a Go program.
// But before this runs, ALL init() functions in the package execute first.
func main() {
	fmt.Println("Hello from main() function")

	// Notice that 'a' was changed inside init(), so now it prints 20.
	fmt.Println("Value of a inside main():", a)
}

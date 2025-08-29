package main

import "fmt"

// This is a global variable 'a' declared outside of any function.
// Its scope is the entire package (in this case, main).
var a = 10

func main() {
	// 'age' is a local variable to the main function.
	age := 30

	// Simple if condition
	if age > 18 {
		// Here we declare a new variable 'a' with ':='.
		// This does NOT change the global 'a' defined above.
		// Instead, it creates a new local variable 'a' inside this 'if' block.
		// This is called VARIABLE SHADOWING because the new 'a'
		// "shadows" or hides access to the global 'a' inside this block.
		a := 47

		// This will print the local 'a' (value 47), not the global one.
		fmt.Println(a) // Output: 47
	}

	// Outside the 'if' block, the local 'a' goes out of scope.
	// Now we can only access the global 'a' again, which is still 10.
	fmt.Println(a) // Output: 10
}

package main

import "fmt"

func main() {
	/*
		int, int8, int16, int32, in64
		float32, float64
		bool
		string
	*/

	/*
		variablename := value

		 In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

		 Can only be used inside functions
	*/

	a := 10
	b := "Hello world"
	c := true
	d := 10.42

	/* a = true
	This wont work as 'a' is already declared as an int
	*/

	// When we update a preassigned value we dont need :=, only =
	a = 11
	b = "Hello Hell"
	c = false
	d = 42.10

	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)
	fmt.Println("Value of c:", c)
	fmt.Println("Value of d:", d)

	/*
		var variablename type = value
		You always have to specify either type or value (or both).

		Can be used inside and outside of functions
	*/
	var x int = 10
	var y string = "GO GO GO"
	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)

	/*
		const CONSTNAME type = value
		The value of a constant must be assigned when you declare it.
		can be declared both inside and outside of a function
	*/
	const Z int = 11 // Typed constant
	const W = true   // Untyped constant
	fmt.Println("Value of z:", Z)
	fmt.Println("Value of w:", W)
}

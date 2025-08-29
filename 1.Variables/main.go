package main

import "fmt"

func main() {
	/*
		Go supports basic data types:

		Integers: int, int8, int16, int32, int64
		Floats:   float32, float64
		Boolean:  bool
		Strings:  string
	*/

	/*
		Short Variable Declaration (:=)
		--------------------------------
		- Syntax: variablename := value
		- Type is INFERRED from the value (compiler decides).
		- Can ONLY be used INSIDE functions.
	*/

	a := 10          // int
	b := "Hello Go!" // string
	c := true        // bool
	d := 10.42       // float64

	/*
		a = true  // ❌ invalid → 'a' is already an int
	*/

	// Updating existing variables → use "=" (not :=)
	a = 11
	b = "Hello World"
	c = false
	d = 42.10

	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)
	fmt.Println("Value of c:", c)
	fmt.Println("Value of d:", d)

	/*
		Standard Variable Declaration (var)
		-----------------------------------
		- Syntax: var variablename type = value
		- Type or value (or both) must be given.
		- Can be used both inside AND outside of functions.
	*/

	var x int = 10
	var y string = "GO GO GO"
	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)

	/*
		Constants (const)
		-----------------
		- Syntax: const CONSTNAME type = value
		- Value must be assigned at declaration (cannot be left uninitialized).
		- Can be declared inside OR outside functions.
		- Cannot be changed later.
	*/

	const Z int = 11 // Typed constant
	const W = true   // Untyped constant
	fmt.Println("Value of Z:", Z)
	fmt.Println("Value of W:", W)
}

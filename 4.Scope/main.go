package main

import (
	"fmt"

	"example.com/mathlib"
)

var (
	a = 30
	b = 20
)

// The add function is now package scoped
// func add(x int, y int) {
// 	z := x + y
// 	fmt.Println(z)
// }

/*
1. block -> { }
*/

func main() {
	var p = 40
	var q = 50
	// The add function is now package scoped
	fmt.Println("Using custom package")
	mathlib.Add(p, q)
	mathlib.Add(a, b)
	mathlib.Sum()
	fmt.Println(mathlib.Money)
	// add (z, a) // Doesnt work as z onky lives within add
}

package main

import "fmt"

var a = 30
var b = 20

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	var p = 40
	var q = 50
	add(p, q)
	add(a, b)
	add(a, p)
	// add (z, a) // Doesnt work as z onky lives within add
}

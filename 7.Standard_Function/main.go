package main

import "fmt"

// 1. Standard or Named Function
// A normal function with a name that can be called multiple times.
func add(x int, y int) {
	fmt.Println("Standard Function:", x+y)
}

func main() {
	add(10, 12)
}

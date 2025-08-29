package main

import "fmt"

var (
	a = 10
	b = 20
)

func printNum(num int) {
	fmt.Println(num)
}

func add(x int, y int) int {
	sum := x + y
	printNum(sum)
	return sum
}

// Below is the main() function
func main() {
	// everything inside is block/local scoped
	add(a, b) //this is coming from global scope as it is not found within local scope
}

package main

import "fmt"

// Variadic function
func print(numbers ...int) { // here numbers is the name of the parameter and it takes the number as a slice and then works with it, we use it to pass unknown number of elements into a slice
	fmt.Println("emp: ", numbers, "len: ", len(numbers), "cap: ", cap(numbers))
}

func main() {
	print(5, 6, 7, 8, 9)
}

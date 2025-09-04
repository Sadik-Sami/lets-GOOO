package main

import "fmt"

var (
	arr3 = [3]string{"Hello", "Mallow", "Shallow"}
)

func main() {
	var arr [2]int
	arr[1] = 6 // we assign 6 to arr's index 1
	arr[0] = 2 // we assign 2 to arr's index 0

	// variable name := [length] type{values of each index}
	arr2 := [2]int{1, 5}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

/*
2 Phases:
	1. Compilation Phase (Compile time)
	2. Execution Phase (Run time)

	***** Compile Segment *****
	main = funct () { ... }

	go run main.go => compile it => main => ./main
	go buid main.go => compile it => main
	
	./main
*/

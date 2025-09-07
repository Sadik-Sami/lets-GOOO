package main

import "fmt"

// func main() {
// 	var x []int      // [], len = 0, cap = 0
// 	x = append(x, 1) // [1], len = 1, cap = 1
// 	x = append(x, 2) // [1, 2], len = 2, cap = 2
// 	x = append(x, 3) // [1, 2, 3], len = 3, cap = 4

// 	y := x

// 	x = append(x, 4)
// 	y = append(y, 5)

// 	x[0] = 10

// 	fmt.Println("emp: ", x, "len: ", len(x), "cap: ", cap(x)) // [10, 2, 3, 5]
// 	fmt.Println("emp: ", y, "len: ", len(y), "cap: ", cap(y)) // [10, 2, 3, 5]
// }
func changeSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]
	y := changeSlice(a)

	fmt.Println("emp: ", x, "len: ", len(x), "cap: ", cap(x)) // [1, 2, 3, 4, 10, 6, 7]
	fmt.Println("emp: ", y, "len: ", len(y), "cap: ", cap(y)) // [10, 6, 7, 11]

	fmt.Println(x[0:8]) // [1, 2, 3, 4, 10, 6, 7, 11] we forcefully get all the elements as we have them in the memory
}

/*
slice underlying array rule => 1024 -> 100% increase after that 25% increase
*/

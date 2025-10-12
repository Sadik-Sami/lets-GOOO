package main

import "fmt"

func main() {
	var a int8 = -128
	var b int8 = 127

	var x uint8 = 10 // unsigned (this has no sign)
	fmt.Println(a + b)
	fmt.Println(a, b, x)
}

package main

import "fmt"

func calculate() (result int) { // named return variable 'result'
	fmt.Println("First", result)
	
	show := func() {
		result = result + 10
		fmt.Println("Defer:", result)
	}
	defer show()

	result = 5

	p := func(a int) {
		fmt.Println("Inside p:", a)
	}
	defer p(result)

	defer fmt.Println(result)

	fmt.Println("second", result)

	defer fmt.Println("Last defer:", 5)
	return
}

func main() {
	a := calculate()
	fmt.Println("main first:", a)
}

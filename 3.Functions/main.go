package main

import "fmt"

/*
func FunctionName(params type) {
  code to be executed
}
*/

func add(num1 int, num2 int) {
	value := num1 + num2
	fmt.Println(value)
}

/*
func FunctionName(params type) type {
  code to be executed
	return
}
*/
func multiply(num1 int, num2 int) int {
	value := num1 * num2
	return value
}

/*
func FunctionName(params type) (type, type) {
  code to be executed
	return
}
*/
func add_multiply(num1 int, num2 int) (int, int) {
	add := num1 + num2
	multiplied := multiply(num1, num2)

	return add, multiplied
}

func welcomeMessage() {
	fmt.Println("Welcome to the application")
}

func getUserName() string {
	// name := '' //This also works as we are starting with an empty string
	var name string
	fmt.Println("Enter your name -")
	fmt.Scanln(&name)
	return name
}

func getNumbers() (int, int) {
	var x, y int
	fmt.Println("Enter first number:")
	fmt.Scanln(&x)
	fmt.Println("Enter second number:")
	fmt.Scanln(&y)
	return x, y
}

func display(name string, sum int, mul int) {
	println("Hello, ", name)
	println("Summation: ", sum)
	println("Multiplication: ", mul)
}

func goodByeMessage() {
	fmt.Println("Thank you for using the application")
	fmt.Println("GoodBye")
}

func main() {
	welcomeMessage()
	name := getUserName()
	num1, num2 := getNumbers()
	sum, mul := add_multiply(num1, num2)
	display(name, sum, mul)
	goodByeMessage()
}

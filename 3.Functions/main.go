package main

import "fmt"

/*
Basic Function Syntax
---------------------
func FunctionName(params type) {
    code to be executed
}
*/

// Function with no return
func add(num1 int, num2 int) {
	value := num1 + num2
	fmt.Println("Addition:", value)
}

/*
Function with Single Return
---------------------------
func FunctionName(params type) returnType {
    code
    return value
}
*/
func multiply(num1 int, num2 int) int {
	value := num1 * num2
	return value
}

/*
Function with Multiple Returns
------------------------------
func FunctionName(params type) (type, type) {
    code
    return value1, value2
}
*/
func add_multiply(num1 int, num2 int) (int, int) {
	added := num1 + num2
	multiplied := multiply(num1, num2)
	return added, multiplied
}

// Function with no params and no return
func welcomeMessage() {
	fmt.Println("Welcome to the application")
}

// Function with no params but WITH return
func getUserName() string {
	var name string // Default zero-value string = ""
	fmt.Println("Enter your name -")
	fmt.Scanln(&name) // Take input from user
	return name
}

// Function with multiple returns (input from user)
func getNumbers() (int, int) {
	var x, y int
	fmt.Println("Enter first number:")
	fmt.Scanln(&x)
	fmt.Println("Enter second number:")
	fmt.Scanln(&y)
	return x, y
}

// Function with params, no return
func display(name string, sum int, mul int) {
	fmt.Println("Hello,", name)
	fmt.Println("Summation:", sum)
	fmt.Println("Multiplication:", mul)
}

// Function with no params, no return
func goodByeMessage() {
	fmt.Println("Thank you for using the application")
	fmt.Println("Goodbye")
}

func main() {
	// Welcome user
	welcomeMessage()

	// Get username
	name := getUserName()

	// Get two numbers from user
	num1, num2 := getNumbers()

	// Call function with multiple return values
	sum, mul := add_multiply(num1, num2)

	// Display the results
	display(name, sum, mul)

	// Exit message
	goodByeMessage()
}

package main

import "fmt"

func main() {
	age := 20
	if age > 18 {
		fmt.Println("You are eligible to be married")
	} else if age < 18 {
		fmt.Print("You are not eligible to be married but you can love someone")
	} else if age == 18 {
		fmt.Println("You are just a teenager, not eligible to be married")
	} else {
		fmt.Println("You are an Alien")
	}

	age = 18
	if age >= 18 {
		fmt.Println("You can vote")
	} else if age < 18 {
		fmt.Println("No voting for you")
	} else {
		fmt.Println("You are an alien")
	}

	/*
		> , >= , <, <=, ==
		and = &&
		or = ||
		not = !
	*/

	age = 25
	sex := "Male"
	if age >= 18 && sex == "Male" {
		fmt.Println("You are eligible for military service")
	} else {
		fmt.Println("You are not eligible for military service")
	}

	isPretty := true
	if isPretty {
		fmt.Println("You are eligible for dating")
	}

	/*
			switch expression {
		case value1:
		    code runs if expression == value1
		case value2:
		    code runs if expression == value2
		default:
		    code runs if none of the above match
		}
	*/

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}

	mark := 78
	switch {
	case mark >= 90:
		fmt.Println("Grade: A, GPA: 4.00")
	case mark >= 85:
		fmt.Println("Grade: A-, GPA: 3.70")
	case mark >= 80:
		fmt.Println("Grade: B+, GPA: 3.30")
	case mark >= 75:
		fmt.Println("Grade: B, GPA: 3.00")
	case mark >= 70:
		fmt.Println("Grade: B-, GPA: 2.70")
	case mark >= 65:
		fmt.Println("Grade: C+, GPA: 2.30")
	case mark >= 60:
		fmt.Println("Grade: C, GPA: 2.00")
	case mark >= 55:
		fmt.Println("Grade: C-, GPA: 1.70")
	case mark >= 50:
		fmt.Println("Grade: D, GPA: 1.00")
	default:
		fmt.Println("Grade: F, GPA: 0.00")
	}
}

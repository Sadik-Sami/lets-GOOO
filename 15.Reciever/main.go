package main

import "fmt"

// -------------------------------
// THEORY: STRUCTS IN GO
// -------------------------------
/*
A struct is a **user-defined data type** in Go.
   - It groups multiple fields (variables) under a single name.
   - Each field has a name and a type.
   - Structs are like "blueprints" for objects.

⚡ Example:
   type User struct {
       Name string
       Age  int
   }

Then we can create (instantiate) objects of type User:
   user1 := User{Name: "Sadik", Age: 30}
   user2 := User{Name: "Al Sami", Age: 32}

Structs in Go are the foundation of building custom types
and are often used instead of "classes" (since Go has no classes).
*/

// Struct declaration
type User struct {
	Name string // field (property) of struct
	Age  int    // another field
}

// -------------------------------
// Normal Function (Parameter Style)
// -------------------------------
/*
This is just a regular function.
It takes a User struct as a parameter.
Notice that this is not "attached" to User,
it just happens to receive a User as an argument.
*/
func printUserDetails(usr User) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// -------------------------------
// Receiver Function (Method)
// -------------------------------
/*
In Go, we can "attach" functions to types using receivers.
This makes them behave like "methods" of that type.

Syntax:
    func (receiverName ReceiverType) methodName(...) { ... }

There are 2 types:
   - Value Receiver (works on a copy of the struct, cannot modify original)
   - Pointer Receiver (works on original struct, can modify fields)
*/

// Value Receiver example (READ-ONLY)
func (u User) printDetails() {
	fmt.Println("Receiver Method -> Name:", u.Name)
	fmt.Println("Receiver Method -> Age:", u.Age)
}

func (u User) call(a int) {
	fmt.Println("Name: ", u.Name)
	fmt.Println("input value is: ", a)
}

// Pointer Receiver example (CAN MODIFY ORIGINAL)
func (u *User) birthday() {
	u.Age++ // increases Age in the original struct (not a copy!)
	fmt.Println("Happy Birthday,", u.Name, "🎉 Now Age:", u.Age)
}

// -------------------------------
// main function
// -------------------------------
func main() {
	// Instantiating structs (objects of type User)
	user1 := User{
		Name: "Sadik",
		Age:  30,
	}
	user2 := User{
		Name: "Al Sami",
		Age:  32,
	}

	// Call a normal function
	printUserDetails(user1)
	printUserDetails(user2)

	// Call value receiver method
	user1.printDetails()
	user2.printDetails()
	user1.call(5)

	// Call pointer receiver method
	user1.birthday()
	user1.birthday()
}

// -------------------------------
// COMPILATION & EXECUTION PHASE
// -------------------------------
/*
When you run: go run main.go
Go does 2 main things:

1) COMPILATION PHASE (compile time)
   - The Go compiler parses and type-checks the code.
   - It builds an intermediate representation of:
       ** Code Segment **
           User = type User struct { Name string; Age int }
           printUserDetails = func(u User) { ... }
           (User).printDetails = func(u User) { ... } // User type object can only call this
					 (User).call = func(u User) (a int) { ... } // User type and takes argument a
					 (User).printDetails = func(u User) {...}
           (*User).birthday = func(*User) { ... }
           main = func() { ... }

   - Struct types and functions are "registered" in the code segment.
   - Constants and global variables are stored in the Data Segment.
   - The compiler performs **escape analysis** to see which variables
     should go on the stack or heap.

   👉 Output of this step is a binary executable file (machine code).

   Example:
       go build main.go → produces ./main (binary executable)

2) EXECUTION PHASE (runtime)
   - Program starts from main()
   - Stack is created for main()
   - Objects (like user1, user2) are created on the stack.
   - Functions are called:
        - printUserDetails(user1) passes a COPY of user1
        - user1.showDetails() runs with a value receiver
        - user1.birthday() runs with a pointer receiver and modifies original
   - Garbage Collector manages memory (cleans heap allocations if any).

   Final Output:
       Name: Sadik
       Age: 30
       Name: Al Sami
       Age: 32
       Receiver Method -> Name: Sadik
       Receiver Method -> Age: 30
       Receiver Method -> Name: Al Sami
       Receiver Method -> Age: 32
       Happy Birthday, Sadik 🎉 Now Age: 31
       Happy Birthday, Sadik 🎉 Now Age: 32

Value Receiver Methods:
  - Operate on a copy of the struct.
  - Cannot modify original struct fields.
  - Can access fields and perform computations.
  - Can accept additional parameters (like call(a int)).

Pointer Receiver Methods:
  - Operate on the original struct.
  - Can modify fields (like birthday()).

*/

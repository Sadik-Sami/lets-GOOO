package main

import "fmt"

// ---------------------------
// Struct Definition
// ---------------------------
// A struct is a collection of fields (variables) grouped together
// - It's like a "custom data type"
// - Similar to "classes" in OOP languages (but without inheritance)
// - Each field has a name + type
type User struct {
	Name string // field: Name of the user (string type)
	Age  int    // field: Age of the user (int type)
}

// ---------------------------
// main() function (entry point)
// ---------------------------
func main() {
	// ---------------------------
	// 1. Declaring and Initializing Struct (Method 1)
	// ---------------------------
	// Create a variable user1 of type User
	var user1 User

	// Assign values to fields (using struct literal)
	user1 = User{
		Name: "Sadik",
		Age:  30,
	}

	// Access fields using dot (.)
	fmt.Println("Name:", user1.Name)
	fmt.Println("Age:", user1.Age)

	// ---------------------------
	// 2. Declaring and Initializing Struct (Method 2: shorthand)
	// ---------------------------
	// Using := shorthand for cleaner initialization
	user2 := User{
		Name: "Al Sami",
		Age:  32,
	}

	fmt.Println("Name:", user2.Name)
	fmt.Println("Age:", user2.Age)

	// ---------------------------
	// 3. Updating Struct Field
	// ---------------------------
	user2.Age = 33 // updating the Age field
	fmt.Println("Updated Age of user2:", user2.Age)

	// ---------------------------
	// 4. Zero Values in Struct
	// ---------------------------
	// If you don't assign values, fields take the "zero value"
	// - string → ""
	// - int → 0
	// - bool → false
	var user3 User
	fmt.Println("User3 (zero values):", user3.Name, user3.Age)

	// ---------------------------
	// 5. Struct Literals without Field Names
	// ---------------------------
	// Less safe (depends on field order)
	// User{Name, Age}
	user4 := User{"Rahim", 25}
	fmt.Println("User4:", user4.Name, user4.Age)
}

/*
===============================
📘 THEORY & NOTES: STRUCTS IN GO
===============================

🔹 What is a Struct?
- A struct is a way to group related data into one unit.
- Think of it as a "record" or "blueprint" for objects.
- Unlike classes in OOP, structs don't have methods directly,
  but you can define methods on structs (via receivers).

🔹 Where does a struct live in memory?
- Struct TYPE (definition) → stored in CODE SEGMENT.
- Struct VARIABLES (instances like user1, user2) → stored in STACK/HEAP depending on usage.

🔹 Why Structs?
- To represent real-world entities (e.g., User, Product, Student).
- To bundle multiple values together instead of using separate variables.

🔹 Struct Initialization:
1. Declare variable then assign → var u User; u = User{Name:"X", Age:20}
2. Short-hand → u := User{Name:"X", Age:20}
3. Without field names → u := User{"X", 20}  (not recommended)
4. Zero value → var u User → u.Name == "" and u.Age == 0

🔹 Accessing & Modifying Fields:
- Dot (.) notation → u.Name, u.Age
- Fields can be updated (structs are mutable).

🔹 Key Difference from OOP:
- Go structs don't have inheritance (no parent/child classes).
- But Go has "composition" → one struct can embed another.
- Methods can be attached to structs using "receiver functions".

---------------------------------
⚙️ COMPILATION PHASE (what Go does)
---------------------------------
- Compiler sees: type User struct {...} → stored in CODE SEGMENT
- Compiler sees: func main() {...} → stored in CODE SEGMENT
- Global constants/variables → stored in DATA SEGMENT
- "go build main.go" → compiles to binary "main"
- "go run main.go" → compiles + runs binary immediately

---------------------------------
⚡ EXECUTION PHASE (when program runs)
---------------------------------
- main() executes.
- user1, user2, etc. are created in STACK (or HEAP if escape analysis decides).
- Struct fields are filled with values.
- Garbage collector (GC) will clean heap values when unused.
*/

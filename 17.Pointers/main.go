package main

import "fmt"

// ----------------------------
// THEORY QUICK NOTES
// ----------------------------
/*
1. MEMORY SEGMENTS in Go:
   - Code Segment → where compiled machine instructions live.
   - Data Segment → global + package-level variables & constants.
   - Stack → stores local variables, function calls, parameters.
   - Heap → dynamic allocations (when data must "escape" function scope).
     Go’s compiler decides (via escape analysis) whether something
     lives on the stack or heap.

2. POINTERS:
   - `&` → "address of" (where a variable is stored in memory).
   - `*` → "value at address" (dereference the pointer).

3. PASS BY VALUE vs PASS BY REFERENCE:
   - Go always passes function arguments **by value** (a copy).
   - If you pass a pointer (`*T`), the copy is just the address,
     meaning the function can access/modify the original object.
   - Example:
       func f(x int) { ... }   // gets a copy of int
       func f(x *int) { ... }  // gets a copy of address → can modify original

4. ARRAYS vs SLICES:
   - Arrays are fixed size and passed as full copies by default.
   - Slices are references to an underlying array, so they "act like" pointers.
*/

// ----------------------------
// Struct Example
// ----------------------------
type User struct {
	Name     string
	Age      int
	Salary   float64
	FavFoods []string
}

// Pass by reference (pointer to array)
// Notice how we use *[3]int instead of [3]int
// - If we used [3]int → full array gets copied into function stack frame.
// - With *[3]int → only address is copied, but function can see/modify original.
func printNumbers(numbers *[3]int) {
	fmt.Println("input (pointer itself): ", numbers)              // shows pointer value (address of arr[0])
	fmt.Println("Address of the Array (pointer var): ", &numbers) // address of pointer var (on stack)
	fmt.Println("Value at pointer (actual array): ", *numbers)    // dereference → original array
}

func main() {
	// ----------------------------
	// Pointers with simple int
	// ----------------------------
	x := 20    // stored on stack (inside main’s stack frame)
	addr := &x // addr is a pointer (stack variable holding x’s address)

	fmt.Println("X:", x)                        // prints 20
	fmt.Println("Address of X:", addr)          // prints memory address
	fmt.Println("Value of X via *addr:", *addr) // dereference → 20

	x = 30     // normal update
	*addr = 30 // update using pointer (affects original x)

	fmt.Println("Value after updating with *addr:", x) // 30

	// ----------------------------
	// Arrays → copied vs referenced
	// ----------------------------
	arr := [3]int{1, 2, 3}

	// printNumbers(arr)  // ❌ invalid, expects pointer
	printNumbers(&arr) // ✅ pass address of array → avoids copy

	// ----------------------------
	// Structs and pointers
	// ----------------------------
	sami := User{
		Name:   "Simanto",
		Age:    32,
		Salary: 35000,
	}
	addrSami := &sami // address of struct (on stack or heap depending on escape analysis)

	fmt.Println("Struct value:", sami)
	fmt.Println("Address of Struct:", addrSami)
	fmt.Println("Dereference Struct pointer:", *addrSami)

	// Shortcut in Go → pointer automatically dereferenced for field access
	fmt.Println("Access field via pointer:", addrSami.Name) // no need for (*addrSami).Name
}

/*
----------------------------
UNDER THE HOOD SUMMARY
----------------------------

1) Compilation Phase:
   - Compiler stores compiled code in **Code Segment**.
   - Global/package-level vars/constants go in **Data Segment**.
   - Compiler does escape analysis:
        If a variable is returned or used outside its scope → allocate on heap.
        Else → allocate on stack.

2) Execution Phase:
   - main() gets a stack frame.
   - `x`, `arr`, and `sami` allocated in stack frame of main (unless they escape).
   - `&x`, `&arr`, `&sami` are pointers holding addresses of these objects.
   - Function calls:
        - If pass `[3]int` → new copy on callee’s stack frame.
        - If pass `*[3]int` → copy of pointer (address), not the array.

3) Garbage Collector:
   - If something goes to the heap, Go’s garbage collector later frees it
     when no references remain.
*/

package main

import "fmt"

// Global variable -> stored in the Data Segment (Global Memory)
var (
	a = 10
)

// Named function -> goes into the Code Segment
func add(x int, y int) {
	// x, y, and z are local variables
	// These are stored in the Stack (temporary memory for each function call)
	z := x + y
	fmt.Println(z)
}

// init() runs automatically before main()
// This is also stored in Code Segment
func init() {
	fmt.Println("Hello Go")
}

func main() {
	// function calls -> new Stack Frame created for each call
	add(5, 4) // arguments stored in stack frame
	add(a, 3) // here `a` is taken from Data Segment (global memory)
}

/*
-------------------------------
 Internal Memory in Go (Simplified)
-------------------------------

1. Code Segment (Text Segment)
   - Stores compiled instructions (all your functions: add, init, main).
   - Functions themselves live here.

2. Data Segment (Global Memory)
   - Stores global & package-level variables (`a = 10`).
   - Exists for the lifetime of the program.

3. Stack
   - Each time a function is called, Go creates a "stack frame".
   - Local variables (x, y, z in `add`) live here temporarily.
   - When the function ends, the stack frame is destroyed (memory freed automatically).

4. Heap
   - Used when variables need to "live beyond" the function call.
   - Example: when you allocate memory dynamically (with `new` or `make`).
   - Garbage Collector (GC) automatically cleans up unused heap memory.

-------------------------------
 What happens in THIS program
-------------------------------

- Step 1: Program starts → `init()` runs → prints "Hello Go".
- Step 2: `main()` starts.
- Step 3: First call: `add(5, 4)`
      -> x=5, y=4 (stack)
      -> z=9 (stack)
      -> prints 9
      -> stack frame destroyed (z is gone).
- Step 4: Second call: `add(a, 3)`
      -> `a` is fetched from Data Segment (global memory = 10)
      -> x=10, y=3 (stack)
      -> z=13 (stack)
      -> prints 13
      -> stack frame destroyed.

Throughout the program:
- `a` always stays in Data Segment.
- Functions (`main`, `add`, `init`) are in Code Segment.
- Stack is created/destroyed on each function call.
- Heap is not used in this example (but GC would watch it if we used it).
*/

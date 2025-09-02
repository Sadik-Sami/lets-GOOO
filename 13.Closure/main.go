package main

import "fmt"

// -------------------------------
// Global variables & constants
// -------------------------------
const a = 10 // "a" is a constant, stored in code segment (doesn't change)
var (
	p = 100 // "p" is a global variable, stored in data segment
)

// -------------------------------
// outer() is a FIRST-CLASS function
// Meaning: It returns another function (a closure)
// -------------------------------
func outer() func() {
	money := 100 // local variable (normally goes to stack)
	age := 30
	fmt.Println("Age: ", age)

	// show() is an anonymous function (closure)
	// It *captures* variables from outer(), especially "money".
	show := func() {
		// "money" is modified here
		// Problem: money belongs to outer(), which should die after outer() ends
		// Solution: Go does "escape analysis" → moves "money" to HEAP
		money = money + a + p
		fmt.Println(money)
	}

	// Returning "show" function
	// This makes "money" live on heap (so it survives after outer() ends)
	return show
}

// -------------------------------
// call() demonstrates closures
// -------------------------------
func call() {
	// increment1 gets a closure from outer()
	increment1 := outer()
	increment1() // First call, money starts at 100 + a + p
	increment1() // Second call, same "money" keeps increasing (closure keeps state alive!)

	// Another closure instance
	increment2 := outer()
	increment2() // Fresh money = 100 again
	increment2() // Increments independently from increment1
}

// -------------------------------
// init() runs automatically before main()
// -------------------------------
func init() {
	fmt.Println("### Bank ###")
}

// -------------------------------
// main()
// -------------------------------
func main() {
	call()
}

/*
=====================================
NOTES: Closure + Memory in Go
=====================================

🔹 Phases in Go
1. Compilation Phase (compile time)
   - Compiler checks syntax, optimizes, and does "escape analysis"
   - Example: It sees that "money" will be used outside "outer()"
     → so it "escapes" the stack and is stored in HEAP.

   Stored in Code Segment:
     - outer = func(){...}
     - outerAnonymous1 (the "show" function) = func(){...}
     - call = func(){...}
     - main = func(){...}
     - init = func(){...}
   Stored in Data Segment:
     - const a = 10
     - var p = 100

2. Execution Phase (runtime)
   - Program actually runs (stack, heap allocated dynamically)

🔹 Closure in Action
- When we call outer(), "money" = 100 is created
- Normally, money would vanish when outer() finishes (stack frame cleared)
- BUT since "show" uses "money", Go moves "money" → HEAP (escape analysis)
- "show" is returned with a reference to "money"
- Now, even after outer() ends, "money" is alive, bound to that closure

🔹 Multiple Closures
- increment1 := outer() → creates its own "money" (100 in HEAP)
- increment2 := outer() → creates a *new* "money" (another 100 in HEAP)
- Both closures are independent, because each "outer()" call makes a new heap allocation

🔹 Garbage Collector
- GC watches the HEAP
- When no closures reference "money" anymore → GC will clean it up

=====================================
Example Execution:
=====================================

### Bank ###     <-- from init()

Age: 30         <-- from first outer()
115             <-- 100 + 10 + 100
225             <-- (115 + 10 + 100)

Age: 30         <-- from second outer()
115             <-- new money = 100 + 10 + 100
225             <-- independent closure, starts fresh from 100

=====================================

So closures = FUNCTION + CAPTURED VARIABLES
Captured variables (like money) escape the stack → heap
Heap is managed by Garbage Collector

*/

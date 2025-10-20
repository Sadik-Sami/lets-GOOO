package main

import (
	"fmt"
)

/*
====================================================
🔹 UNDERSTANDING DEFER — COMPLETE EXPLANATION
====================================================

💡 Concept:
"defer" schedules a function call to run *after* the
current function finishes execution — either normally
or during panic (unwinding).

When Go sees:
    defer fmt.Println("value:", i)

➡ Step 1: Go *evaluates* "fmt.Println" and its arguments ("i")
           immediately at that line.
➡ Step 2: It stores a "defer record" — function pointer + arguments
           — in a small list attached to the current goroutine’s stack.
➡ Step 3: When the function returns, all deferred calls are executed
           in reverse order (LIFO: Last-In-First-Out).

====================================================
🔹 MEMORY VIEW (Simplified 32-bit layout)
====================================================

| Code Segment |  compiled code (main, a, fmt.Printf, etc.)
| Data Segment |  global constants, static variables
| Stack        |  function frames, locals, defer records
| Heap         |  dynamically allocated memory (strings, slices, etc.)

When a function is called:
    - CPU allocates a new "stack frame"
    - Locals (like i) live inside that frame
    - Any defer statements create small records on stack
====================================================
*/

func a() {
	/*
		At this point, the runtime allocates a new stack frame
		for "a" inside the goroutine’s stack.
	*/
	i := 0                                 // local variable 'i' stored on stack
	fmt.Printf("First execution: %d\n", i) // prints 0

	// ❗️Important: i is evaluated *now* (value 0)
	// A "defer record" is created containing:
	//    fn pointer -> fmt.Printf
	//    arguments  -> ("defer's execution: %d\n", 0)
	// This record is pushed onto the goroutine's internal defer stack.
	defer fmt.Printf("defer's execution: %d\n", i)

	i++                                    // now i = 1 (changes local var, not defer arg)
	fmt.Printf("Third execution: %d\n", i) // prints 1

	/*
		When function "a" returns:
			- Go runtime checks for defers in this frame
			- Pops the last defer record
			- Executes fmt.Printf("defer's execution: 0")
		Because arguments were captured at defer time (value = 0)
	*/
}

func main() {
	a()

	/*
		Expected Output:
		----------------
		First execution: 0
		Third execution: 1
		defer's execution: 0
	*/
	runExamples()
}

/*
====================================================
🔹 WHAT HAPPENS INTERNALLY
====================================================

Let's simulate how Go executes "a()":

1️⃣  Stack frame created for a()
     locals: i = 0
     defer list: empty

2️⃣  i := 0  → stored on stack frame
     fmt.Printf → prints "First execution: 0"

3️⃣  defer fmt.Printf("defer's execution: %d\n", i)
     - evaluated immediately:
          "fmt.Printf" identified
          "i" evaluated → 0
     - runtime pushes defer record (fn pointer, args) onto stack

4️⃣  i++ → 1
     fmt.Printf("Third execution: 1")

5️⃣  return from a()
     - runtime pops defer stack (LIFO)
     - executes stored fn: fmt.Printf("defer's execution: 0")
====================================================
🔹 DEFER + CLOSURE EXAMPLES
====================================================
*/

// Example showing difference between evaluated argument vs closure capture
func exampleValueDefer() {
	i := 0
	defer fmt.Println("Value defer (evaluated now):", i) // i=0 now
	i++
	fmt.Println("Current i:", i) // i=1
}
func exampleClosureDefer() {
	i := 0
	defer func() {
		fmt.Println("Closure defer (captured variable):", i)
	}() // i is referenced, not evaluated yet
	i++
	fmt.Println("Current i:", i)
}

// Example: multiple defers = LIFO order
func exampleMultipleDefers() {
	defer fmt.Println("First defer (will print last)")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer (will print first)")
}

// Example: defer + panic = executes before crash
func examplePanicDefer() {
	defer fmt.Println("Cleanup defer: runs even on panic")
	panic("💥 something went wrong")
}

// Example: defer + recover = gracefully handle panic
func exampleRecoverDefer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	panic("🔥 runtime error but recovered!")
}

/*
====================================================
🔹 INTERVIEW NOTES
====================================================

Q1. When are defer arguments evaluated?
    ➤ Immediately when defer statement is executed.

Q2. When are deferred functions executed?
    ➤ When the surrounding function returns
       (on normal return or during panic unwinding).

Q3. What order are multiple defers executed in?
    ➤ LIFO (Last In, First Out).

Q4. Where are defer records stored?
    ➤ In runtime metadata for the goroutine,
       often as a linked list or small stack
       stored on the goroutine’s stack frame.

Q5. Do deferred calls run if a panic happens?
    ➤ Yes. Always. Even during panic.

Q6. Can defer slow down my program?
    ➤ A bit, yes. Each defer adds runtime bookkeeping.
       Avoid defers in *hot loops*, but they’re fine for cleanup.

Q7. Common use-cases?
    ➤ Closing files, releasing locks, cancelling contexts,
       recovering from panics, resource cleanup.

====================================================
🔹 VISUAL STACK TRACE FOR FUNC a()
====================================================

           (Memory grows downward)
           ┌────────────────────────┐
           │ runtime, fmt, globals  │ ← Code/Data
           ├────────────────────────┤
           │ stack frame: main()    │
           │  return address         │
           │  call a()               │
           ├────────────────────────┤
           │ stack frame: a()       │
           │  local var i = 0        │
           │  defer record { fn=fmt.Printf, arg=0 }  │
           │  return address         │
           ├────────────────────────┤
           │ deferred call executes │ ← fmt.Printf("defer's execution: 0")
           ├────────────────────────┤
           │ stack unwinds          │
           └────────────────────────┘

====================================================
🔹 KEY TAKEAWAY
====================================================

Defer lets you write clean, readable cleanup logic.

✅ Arguments are captured *immediately*.
✅ Deferred calls execute *at the end* (LIFO).
✅ Defers always run — even during panics.
✅ Runtime stores defers per goroutine.
✅ Great for closing files, recovering from panics,
   and unlocking resources safely.

====================================================
*/

func runExamples() {
	fmt.Println("\n--- Value Defer Example ---")
	exampleValueDefer()

	fmt.Println("\n--- Closure Defer Example ---")
	exampleClosureDefer()

	fmt.Println("\n--- Multiple Defers Example ---")
	exampleMultipleDefers()

	fmt.Println("\n--- Panic + Defer Example ---")
	// examplePanicDefer()

	fmt.Println("\n--- Recover + Defer Example ---")
	exampleRecoverDefer()
}

/*
HOW TO USE THIS FILE:

1. Run normally:
       go run defer_notes.go

2. Uncomment "runExamples()" in main() to explore all defer cases.

3. For Interviews, recall:
   - Code segment = compiled code
   - Data segment = constants
   - Stack = locals + defers + return addr
   - Heap = dynamic allocations

4. Common trick question:
   Q: What will this print?
      i := 0
      defer fmt.Println(i)
      i++
      → prints 0 (not 1)
====================================================
*/

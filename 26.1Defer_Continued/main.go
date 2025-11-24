package main

import "fmt"

// Note: comments below contain both compilation-phase and execution-phase notes.

func calc() int { // traditional return (unnamed)
	// Compilation phase:
	//  - 'result' is a local variable of type int, stored in this function's frame.
	//  - 'show' is a variable holding a function value (closure). The compiler will
	//    detect that 'show' refers to 'result' and may perform escape analysis.
	// Execution phase (step-by-step):
	//  - stack frame for calc created. space allocated for 'result' and 'show'.
	//  - result initialized to 0 (zero value for int).
	result := 0
	fmt.Println("first:", result)

	show := func() {
		// This closure *captures the variable* 'result' (the variable, not just its value).
		// It will read/modify the same memory cell as 'result'.
		result = result + 10
		fmt.Println("defer:", result)
	}
	// defer show() -> the deferred call is scheduled now; the function value 'show'
	// (the closure) is evaluated now and stored in the goroutine's defer list.
	// Note: arguments to a deferred call are evaluated at the point of the defer statement.
	defer show()

	// mutate result
	result = 5
	fmt.Println("second", result)

	// Return step (unnamed return expression):
	//  1. Evaluate the return expression `result` → produce value 5.
	//  2. Copy that value into a hidden return-value temporary (call it 'retVal').
	//  3. Execute deferred calls (LIFO) — here show() runs and updates the captured 'result'
	//     variable to 15, and prints "defer: 15".
	//  4. Function returns to caller with the already-copied 'retVal' (5).
	return result
}

func calculate() (result int) { // named return variable 'result'
	// Compilation phase:
	//  - A named return variable `result` becomes part of the function's frame.
	//  - Using a naked 'return' will return the current value of this variable.
	fmt.Println("First:", result)

	show := func() {
		// closure captures the same 'result' variable (named return slot).
		result = result + 10
		fmt.Println("Defer:", result)
	}
	defer show()

	result = 5
	fmt.Println("Second:", result)

	// Naked return:
	//  1. Execute deferred calls (they run before the function returns to the caller).
	//     The deferred show() changes `result` to 15.
	//  2. After defers complete, the function returns the current value of named 'result' (15).
	return
}

func main() {
	a := calc()
	fmt.Println("main first:", a)

	b := calculate()
	fmt.Println("main second:", b)
}

/*
--------------------------------------------------------------------------------
ADDITIONAL TIDY NOTES (compilation, runtime, memory layout, commands, experiments)
--------------------------------------------------------------------------------
1) Key semantic summary
- calc(): unnamed-return -> return expression evaluated & copied before defers run -> caller
  receives copied value (5).
- calculate(): named-return -> defers run before final return and can mutate the named
  return slot -> caller receives mutated value (15).

2) Where things live (short answers)
- closure code (instructions): CODE segment (.text)
- closure value (fn_ptr + env_ptr): STACK if not escaping; HEAP if closure escapes
- captured variables (e.g., result): STACK if not escaping; HEAP if escape analysis requires
- defer records: stored in goroutine's defer list on the goroutine stack (no separate stack)

3) Compilation-phase (what compiler performs)
- Parse → AST → SSA IR
- Closure creation & env representation
- Escape analysis (decide stack vs heap for captured vars)
- Stack layout emission (locals, named return slots, closure structs, defer metadata)
- Generate code: prologue, body, epilogue, and metadata (.gopclntab, GC maps)

4) Execution-phase timeline (concise)
- Function prologue: allocate frame, save registers
- Evaluate defer: create deferRec (copy closure value + args) and push onto goroutine defer list
- Execute body (mutate locals)
- For 'return expr' (calc): evaluate expr → copy to ret slot/register -> run defers -> ret
- For naked 'return' (calculate): run defers (they can mutate named slots) -> return

5) Compact ASCII memory layout (conceptual, x86_64, words = 8 bytes)
-- At time just after defer in calc(), before result=5:
[ higher addr ]
  caller return addr
  saved regs
  hidden ret slot (retVal)         <-- (may be unused until return)
  closure struct 'show'           <-- { fn_ptr -> .text, env_ptr -> &result }
  local 'result' (int) = 0        <-- &result used by closure
  defer list head -> deferRec     <-- contains copy of closure struct
[ lower addr ]

-- After result=5 and after return expr is evaluated (calc):
  hidden ret slot = 5
  closure struct
  local 'result' = 5
  deferRec -> points to closure

-- During deferred call:
  hidden ret slot = 5
  closure struct
  local 'result' = 15   <-- updated by defer
  deferRec removed

Note: function returns with ret slot = 5; frame popped after ret, stack memory freed.

6) Defer record (conceptual structure)
- Stored on goroutine stack
- Contains: fn_ptr or closure value, env pointer(s), evaluated args, link to previous defer
- Executed LIFO during normal return or panic unwinding

7) Escape analysis & commands
- Inspect decisions:
    go version
    go build -gcflags="-m" defer_closure.go
- Typical messages: "x does not escape" or "x escapes to heap"
- To force escape for testing: return the closure from a function and rebuild with -m

8) Assembly & debugging commands (practical)
- View SSA-style compile output & assembly:
    go tool compile -S defer_closure.go > asm.txt
  or build with assembly prints:
    go build -gcflags="-S" defer_closure.go  # prints useful optimization notes
- Disassemble built binary:
    go build -o deferprog defer_closure.go
    go tool objdump -s main.calc deferprog
    go tool objdump -s main.calculate deferprog
- Use Delve for live debugging of stack and defer records (advanced):
    dlv debug

9) Small experiments (exact steps)
A) Run the program and confirm output
   go run defer_closure.go
B) See escape analysis decisions
   go build -gcflags="-m" defer_closure.go
C) Force heap escape test
   - Create a function that returns the closure (see notes) and build with -m
   - Expect message: "x escapes to heap"
D) Inspect assembly for where ret slot is filled and where defer call occurs
   go tool compile -S defer_closure.go > asm.txt

10) Quick interview summary (memorables)
- Defer LIFO, defer args evaluated at defer site, deferred calls run before the frame is popped.
- Unnamed return: return expression evaluated & copied before defers. Named return: defers can change named slots.
- Closure captures variables (the environment), not just snapshot values.
- Escape analysis moves captured locals to heap when closure outlives frame.

11) Reference commands (cheat sheet)
- Run:         go run defer_closure.go
- Build w/ -m:  go build -gcflags="-m" defer_closure.go
- Assembly:     go tool compile -S defer_closure.go > asm.txt
- Disasm:       go build -o deferprog defer_closure.go && go tool objdump -s main.calc deferprog

*/

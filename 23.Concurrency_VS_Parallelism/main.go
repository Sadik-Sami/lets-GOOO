package main

import (
	"fmt"
	"runtime"
	"time"
)

// heavyWork simulates a CPU-intensive task
func heavyWork(name string) {
	start := time.Now() // mark start time for this worker
	count := 0
	for i := 0; i < 1e8; i++ { // 100 million iterations
		count += i
	}
	fmt.Println(name, "done in", time.Since(start))
}

func runWithProcs(procs int) {
	// Limit Go to given number of OS threads
	runtime.GOMAXPROCS(procs)

	start := time.Now()

	// Run 4 heavy tasks as goroutines
	done := make(chan bool, 4)
	for i := 1; i <= 4; i++ {
		go func(id int) {
			heavyWork(fmt.Sprintf("Worker-%d", id))
			done <- true
		}(i)
	}

	// Wait for all workers to finish
	for i := 0; i < 4; i++ {
		<-done
	}

	fmt.Printf("With GOMAXPROCS(%d): total time %v\n\n", procs, time.Since(start))
}

func main() {
	fmt.Println("🔹 Concurrency (1 core, time-slicing)")
	runWithProcs(1) // all 4 goroutines share 1 core

	fmt.Println("🔹 Parallelism (use all CPU cores)")
	runWithProcs(runtime.NumCPU()) // goroutines spread across available cores
}

/* *
What’s Happening Under the Hood?

### Concurrency (1 core) ###

>>>We tell Go runtime: “you can only use 1 OS thread”.

*Two goroutines (Worker-1, Worker-2) are started.

*But only one core executes at a time → the runtime time-slices between them.

*They’ll finish one after another, so the total runtime ≈ sum of both.

### Parallelism (multiple cores) ###

>>>Now the runtime can spread goroutines across all CPU cores.

* Both workers execute at the same time on different cores.

* They finish around the same time, so the total runtime ≈ time of the slowest one (almost half).
*/

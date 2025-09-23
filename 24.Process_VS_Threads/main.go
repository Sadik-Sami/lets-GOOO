package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

// -------------------------
// Goroutine "Thread" Switching
// -------------------------
func goroutineSwitches(n int) time.Duration {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)

	// Channel used to force context switches
	ch := make(chan struct{})

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-ch // wait for signal from other goroutine
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			ch <- struct{}{} // send signal (forces switch)
		}
	}()

	wg.Wait()
	return time.Since(start)
}

// -------------------------
// Process Switching
// -------------------------
func processSwitches(n int) time.Duration {
	start := time.Now()

	for i := 0; i < n; i++ {
		// Spawn a new process (very costly compared to goroutines)
		cmd := exec.Command("true") // lightweight dummy command
		_ = cmd.Run()
	}

	return time.Since(start)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	iterations := 10000

	fmt.Printf("Measuring with %d iterations...\n\n", iterations)

	// Goroutine context switch
	gTime := goroutineSwitches(iterations)
	fmt.Printf("Goroutine switching took: %v\n", gTime)

	// Process context switch
	pTime := processSwitches(iterations)
	fmt.Printf("Process switching took:   %v\n", pTime)

	fmt.Println("\n Goroutines are lightweight. Process creation/switching is orders of magnitude heavier.")
}

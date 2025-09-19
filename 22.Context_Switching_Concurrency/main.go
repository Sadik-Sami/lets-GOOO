package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(name string) {
	for {
		fmt.Println(name, time.Now().Format("15:04:05.000"))
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go worker("Worker-1")
	go worker("Worker-2")
	select {}
}

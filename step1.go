// multithreading
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func cpuHeavy(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting %s\n", name)
	total := 0
	for i := 0; i < 10_000_000; i++ {
		total += i * i
	}
	fmt.Printf("%s complete\n", name)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores
	var wg sync.WaitGroup
	wg.Add(2)

	go cpuHeavy("Goroutine A", &wg)
	go cpuHeavy("Goroutine B", &wg)

	wg.Wait()
}

// Illustrates both threads and processes
package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func runGoroutine(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	total := 0
	for i := 0; i < 10000000; i++ {
		total += i * i
	}
	fmt.Printf("Goroutine %d finished in %.4f seconds\n", id, time.Since(start).Seconds())
}

func runProcess(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command("./worker", fmt.Sprintf("%d", id))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Process %d error: %v\n", id, err)
	}
	fmt.Printf("%s", output)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	numTasks := 4

	fmt.Println("=== Running Multithreading with Goroutines ===")
	var wg1 sync.WaitGroup
	wg1.Add(numTasks)
	for i := 0; i < numTasks; i++ {
		go runGoroutine(i, &wg1)
	}
	wg1.Wait()

	fmt.Println("\n=== Running Multiprocessing with Subprocesses ===")
	var wg2 sync.WaitGroup
	wg2.Add(numTasks)
	for i := 0; i < numTasks; i++ {
		go runProcess(i, &wg2)
	}
	wg2.Wait()
}

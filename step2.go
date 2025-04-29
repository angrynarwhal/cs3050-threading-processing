//multithreading

package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func cpuHeavyTask(id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	start := time.Now()

	// Simulated CPU-heavy computation
	total := 0
	for i := 0; i < 10000000; i++ {
		total += i * i
	}

	duration := time.Since(start)
	results <- fmt.Sprintf("%d,%.4f\n", id, duration.Seconds())
}

func main() {
	numWorkers := 10
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	results := make(chan string, numWorkers)

	fmt.Printf("Running %d goroutines...\n", numWorkers)
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go cpuHeavyTask(i, &wg, results)
	}

	wg.Wait()
	close(results)

	// Write results to CSV
	file, err := os.Create("goroutine_times.csv")
	if err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}
	defer file.Close()

	file.WriteString("task_id,time_sec\n")
	for r := range results {
		file.WriteString(r)
	}

	fmt.Println("Timing data saved to goroutine_times.csv")
}

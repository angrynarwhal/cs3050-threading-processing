// Worker
// Must be compiled:  `go build -o worker worker.go`

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func cpuHeavy() {
	total := 0
	for i := 0; i < 10000000; i++ {
		total += i * i
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./worker <id>")
		return
	}

	// ✅ Fixed: capture and check the conversion
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid ID argument")
		os.Exit(1)
	}

	start := time.Now()
	cpuHeavy()
	duration := time.Since(start)
	fmt.Printf("Process %d finished in %.4f seconds\n", id, duration.Seconds())
}

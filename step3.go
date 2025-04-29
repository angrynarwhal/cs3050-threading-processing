// Shows only threads
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func cpuHeavyTask(id int) {
	start := time.Now()
	total := 0
	for i := 0; i < 10000000; i++ {
		total += i * i
	}
	duration := time.Since(start)
	fmt.Printf("Process %d finished in %.4f seconds\n", id, duration.Seconds())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./worker <id>")
		return
	}
	id, _ := strconv.Atoi(os.Args[1])
	cpuHeavyTask(id)
}

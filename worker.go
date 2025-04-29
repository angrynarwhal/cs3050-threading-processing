// Worker
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
	_ = strconv.Atoi(os.Args[1]) // just to simulate input
	start := time.Now()
	cpuHeavy()
	fmt.Printf("Process done in %.4f seconds\n", time.Since(start).Seconds())
}

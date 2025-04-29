package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func cpuHeavy() float64 {
	start := time.Now()
	total := 0
	for i := 0; i < 10000000; i++ {
		total += i * i
	}
	return time.Since(start).Seconds()
}

func runGoroutines(n int) []float64 {
	var wg sync.WaitGroup
	times := make([]float64, n)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			times[idx] = cpuHeavy()
		}(i)
	}
	wg.Wait()
	return times
}

func runProcesses(n int) []float64 {
	var wg sync.WaitGroup
	times := make([]float64, n)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			start := time.Now()
			cmd := exec.Command("./worker", fmt.Sprintf("%d", idx))
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error running process:", err)
			}
			times[idx] = time.Since(start).Seconds()
		}(i)
	}
	wg.Wait()
	return times
}

func makeBarChart(gTimes, pTimes []float64) error {
	n := len(gTimes)
	values := make(plotter.Values, 2*n)

	for i := 0; i < n; i++ {
		values[i] = gTimes[i]
		values[n+i] = pTimes[i]
	}

	p := plot.New()
	p.Title.Text = "Goroutines vs Subprocesses Execution Time"
	p.Y.Label.Text = "Time (s)"

	barWidth := vg.Points(20)
	barChart, err := plotter.NewBarChart(values, barWidth)
	if err != nil {
		return err
	}
	barChart.LineStyle.Width = vg.Length(0)
	barChart.Color = plotter.DefaultColors[1]

	p.Add(barChart)
	p.NominalX(genLabels(n)...)

	return p.Save(6*vg.Inch, 4*vg.Inch, "goroutines_vs_processes.png")
}

func genLabels(n int) []string {
	var labels []string
	for i := 0; i < n; i++ {
		labels = append(labels, fmt.Sprintf("G%d", i))
	}
	for i := 0; i < n; i++ {
		labels = append(labels, fmt.Sprintf("P%d", i))
	}
	return labels
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	num := 5

	fmt.Println("Running goroutines...")
	goroutineTimes := runGoroutines(num)

	fmt.Println("Running processes...")
	processTimes := runProcesses(num)

	fmt.Println("Creating bar chart...")
	err := makeBarChart(goroutineTimes, processTimes)
	if err != nil {
		fmt.Println("Failed to generate plot:", err)
		return
	}
	fmt.Println("Chart saved as goroutines_vs_processes.png")
}

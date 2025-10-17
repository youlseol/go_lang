package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple goroutine example
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Worker function for goroutines
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}

func main() {
	fmt.Println("=== Simple Goroutine ===")
	go printNumbers()
	time.Sleep(600 * time.Millisecond)

	fmt.Println("\n=== WaitGroup Example ===")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d started\n", n)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d finished\n", n)
		}(i)
	}
	wg.Wait()

	fmt.Println("\n=== Channels Example ===")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}

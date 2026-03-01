package main

import (
	"fmt"
	"sync"
)

/*
workers - goroutines that processes the tasks
*/
func workers(jobs <-chan int, wg *sync.WaitGroup) {
	// 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go processJobs(jobs, i, wg)
	}
}

func processJobs(jobs <-chan int, workerId int, wg *sync.WaitGroup) {
	defer wg.Done()
	// process jobs from channel
	for job := range jobs {
		fmt.Println("Worker", workerId, "processing job", job)
	}
}

func main() {
	var wg sync.WaitGroup
	// create jobs to be processed by workers
	jobs := make(chan int, 10)

	workers(jobs, &wg)

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	// sender closes the channel
	close(jobs)

	// main go routine has to wait for go-routines
	wg.Wait()
}

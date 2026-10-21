package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func worker(ctx context.Context, jobs <-chan int, wg *sync.WaitGroup, completedJobs *int64) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Println("processing", job)
			atomic.AddInt64(completedJobs, 1)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		100*time.Millisecond)
	defer cancel()
	jobs := make(chan int, 20)
	var wg sync.WaitGroup
	var completedJobs int64
	var pushedJobs int64

	numWorkers := 3
	for range numWorkers {
		wg.Add(1)
		go worker(ctx, jobs, &wg, &completedJobs)
	}

	go func() {
		defer close(jobs)
		for i := range 20000 {
			select {
			case <-ctx.Done():
				//stop pushing jobs if the context is dead
				return
			case jobs <- i:
				//job sent
				atomic.AddInt64(&pushedJobs, 1)
			}

		}
	}()

	wg.Wait()

	fmt.Printf("Jobs pushed: %d, jobs completed: %d", pushedJobs, completedJobs)
}

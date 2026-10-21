package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	jobs := make(chan int)
	defer wg.Done()

	for {
		select {
		case job := <-jobs:
			fmt.Println(job)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	before := runtime.NumGoroutine()
	ctx, cancel := context.WithCancel(
		context.Background())

	var wg sync.WaitGroup
	for range 99 {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(100 * time.Millisecond)
	cancel()

	wg.Wait()
	after := runtime.NumGoroutine()

	fmt.Println("before", before)
	fmt.Println("after", after)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(500 * time.Millisecond)
	defer timer.Stop()

	work := make(chan struct{})

	go doWork(work)
	select {
	case <-work:
		fmt.Println("completed")
	case <-timer.C:
		fmt.Println("timeout")
	}

}

func doWork(work chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	work <- struct{}{}
}

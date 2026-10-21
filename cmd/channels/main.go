package main

import (
	"fmt"
	"os"
)

func main() {
	runType := "nil-send"
	//fmt.Printf("%v\n", os.Args)

	if len(os.Args) >= 2 {
		runType = os.Args[1]
	}
	//fmt.Printf("%v\n", runType)
	switch runType {
	case "nil-send":
		nilSend()
	case "closed-receive":
		closedReceive()
	case "double-close":
		doubleClose()
	case "closed-send":
		closedSend()
	case "select":
		selectNoDefault()
	case "select-ready":
		selectReady()
	case "worker":
		testWorker()
	}

}

func testWorker() {
	jobs := make(chan int)
	done := make(chan struct{})
	go worker(jobs, done)
	for i := range 20 {
		jobs <- i
		/* simulate empty recieves after jobs channel closed.
		if i == 12 {
			close(jobs)
		} else if i < 12 {
			jobs <- i
		}
		*/
		if i == 13 {
			signal := struct{}{}
			done <- signal
		}
	}
	//signal := struct{}{}
	//done <- signal

}

func worker(jobs <-chan int, done <-chan struct{}) {
	for {
		select {
		case job, ok := <-jobs:
			if ok {
				fmt.Println(job)
			}

		case <-done:
			for job := range jobs {
				fmt.Println(job)
			}
			return
		}
	}
}

func selectReady() {
	foo := make(chan string, 1)
	bar := make(chan string, 1)

	foo <- "rust"
	bar <- "go"

	select {
	case <-foo:
		fmt.Println("FOO")
	case <-bar:
		fmt.Println("BAR")
	}

}

func selectNoDefault() {
	ch := make(chan int)

	select {
	case v := <-ch:
		fmt.Println("received", v)
	}
}

func selectWithDefault() {
	ch := make(chan int)

	select {
	case v := <-ch:
		fmt.Println("received", v)
	default:
		fmt.Println("not ready")
	}
}

func closedSend() {
	ch := make(chan int, 2)
	ch <- 1911
	close(ch)
	ch <- 1776 //panic send on closed channel

}

func doubleClose() {
	ch := make(chan int, 2)
	ch <- 1911
	ch <- 1776
	close(ch)
	close(ch) //panic close of closed channel

}

func closedReceive() {
	ch := make(chan int, 1)
	ch <- 42
	close(ch)

	v, ok := <-ch
	fmt.Printf("First receive: v=%d, ok:%t\n", v, ok) //42, true

	v, ok = <-ch
	fmt.Printf("Second receive: v:%d, ok:%t\n", v, ok) //0, false

	zed := make(chan int, 3)
	zed <- 13
	zed <- 27
	close(zed) //zed is dead baby, if we don't close the channel we get a deadlock

	for v := range zed {
		fmt.Printf("received: v:%d\n", v)
	}

}

func nilSend() {
	var ch chan int
	ch <- 42 //deadlock
}

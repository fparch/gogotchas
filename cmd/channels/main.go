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
	}
}

func closedSend() {
	
}

func doubleClose() {
	ch := make(chan int, 2)
	ch <- 1911
	ch <- 1776
	close(ch)
	close(ch)
}

func closedReceive() {
	ch := make(chan int, 1)
	ch <- 42
	close(ch)

	v, ok := <-ch
	fmt.Printf("First receive: v=%d, ok:%t\n", v, ok)

	v, ok = <-ch
	fmt.Printf("Second receive: v:%d, ok:%t\n", v, ok)

}

func nilSend() {
	var ch chan int
	ch <- 42
}

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func main() {
	c := Counter{}
	c.Increment()
	fmt.Printf("Counter: %d\n", c.n)
}

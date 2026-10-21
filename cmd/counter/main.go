package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter interface {
	Increment()
	Value() int64
}

type MutexCounter struct {
	mu sync.Mutex
	k  int64
}

func (c *MutexCounter) Increment() {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.k++
}

func (c *MutexCounter) Value() int64 {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.k
}

type RWMutexCounter struct {
	mu sync.RWMutex
	k  int64
}

func (c *RWMutexCounter) Increment() {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.k++
}

func (c *RWMutexCounter) Value() int64 {
	defer c.mu.RUnlock()
	c.mu.RLock()
	return c.k
}

type AtomicCounter struct {
	k atomic.Int64
}

func (c *AtomicCounter) Increment() {
	c.k.Add(1)
}

func (c *AtomicCounter) Value() int64 {
	return c.k.Load()
}

type NaiveCounter struct {
	k int64
}

func (c *NaiveCounter) Increment() {
	c.k++
}

func (c *NaiveCounter) Value() int64 {
	return c.k
}

var (
	_ Counter = (*MutexCounter)(nil)
	_ Counter = (*RWMutexCounter)(nil)
	_ Counter = (*AtomicCounter)(nil)
	_ Counter = (*NaiveCounter)(nil)
)

var Implementations = map[string]func() Counter{
	"mutex":   func() Counter { return &MutexCounter{} },
	"rwmutex": func() Counter { return &RWMutexCounter{} },
	"atomic":  func() Counter { return &AtomicCounter{} },
	"naive":   func() Counter { return &NaiveCounter{} },
}

func RunConcurrent(c Counter, numRoutines, n int) int64 {
	var wg = sync.WaitGroup{}
	for range numRoutines {
		wg.Go(func() {
			for range n {
				c.Increment()
			}
		})
	}
	wg.Wait()
	return c.Value()
}

func main() {
	for _, name := range []string{"mutex", "rwmutex", "atomic"} {
		c := Implementations[name]()
		fmt.Printf("%-8s %d\n", name, RunConcurrent(c, 100, 1000))
	}

	/*
		impls := map[string]func() Counter{
			"mutex":   func() Counter { return &MutexCounter{} },
			"rwmutex": func() Counter { return &RWMutexCounter{} },
			"atomic":  func() Counter { return &AtomicCounter{} },
			"naive":   func() Counter { return &NaiveCounter{} },
		}

		var wg sync.WaitGroup
		counter := impls["mutex"]()

		for range 100 {
			wg.Go(func() {
				for range 1000 {
					counter.Increment()
				}
			})
		}
		wg.Wait()
		fmt.Printf("Counter value: %d", counter.Value())
	*/
}

/*
Your tasks:
- Implement all three.
- Launch 100 goroutines.
- Increment 1,000 times per goroutine.
- Verify the final result is exactly 100,000.
- Run the test with -race.
Then write benchmarks for:
- Single-goroutine increments.
- Parallel increments.
- 90% reads, 10% writes.
*/

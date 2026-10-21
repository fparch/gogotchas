package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			//wg.Add(1)
			defer wg.Done()

			fmt.Println(n)
		}(i)
	}

	wg.Wait()
}

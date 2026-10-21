package main

import (
	"context"
	"fmt"
	"time"
)

func operation() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second)
	defer cancel()

	if ctx.Err() != nil {
		cancel()
	}
	fmt.Println(ctx.Deadline())

}

func main() {
	operation()
}

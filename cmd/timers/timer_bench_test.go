package main

import (
	"testing"
	"time"
)

var sink <-chan time.Time

func BenchmarkTimeAfter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = time.After(time.Hour)
	}
}

func BenchmarkTickerReuse(b *testing.B) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for i := 0; i < b.N; i++ {
		sink = ticker.C
	}
}

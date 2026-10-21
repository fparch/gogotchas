package main

import (
	"sync/atomic"
	"testing"
)

func BenchmarkSingleIncrement(b *testing.B) {
	for name, newC := range Implementations {
		b.Run(name, func(b *testing.B) {
			c := newC()
			b.ResetTimer()
			for range b.N {
				c.Increment()
			}
		})
	}
}

func BenchmarkParallelIncrement(b *testing.B) {
	for name, newC := range Implementations {
		b.Run(name, func(b *testing.B) {
			c := newC()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					c.Increment()
				}
			})
		})
	}
}

// 90% reads, 10% writes, run in parallel.
func BenchmarkMixed90Read10Write(b *testing.B) {
	for name, newC := range Implementations {
		b.Run(name, func(b *testing.B) {
			c := newC()
			b.RunParallel(func(pb *testing.PB) {
				var i uint32
				var sink int64
				for pb.Next() {
					if i%10 == 0 {
						c.Increment()
					} else {
						sink += c.Value()
					}
					i++
				}
				_ = atomic.LoadInt64(&sink) // keep sink live so reads aren't optimized away
			})
		})
	}
}

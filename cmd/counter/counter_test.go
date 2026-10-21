package main

import "testing"

func TestCountersExact(t *testing.T) {
	for _, name := range []string{"mutex", "rwmutex", "atomic"} {
		t.Run(name, func(t *testing.T) {
			got := RunConcurrent(Implementations[name](), 100, 1000)
			if got != 100_000 {
				t.Fatalf("got %d, want 100000", got)
			}
		})
	}
}

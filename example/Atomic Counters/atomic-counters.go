package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The primary mechanism for managing state in Go is communication over channels.
// We saw this for example with worker pools.
// There are a few other options for managing state though. Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
func main() {
	// We’ll use an unsigned integer to represent our (always-positive) counter.
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// In order to safely use the counter while it’s still being updated by other goroutines, we extract a copy of the current value into opsFinal via LoadUint64.
	fmt.Println("ops:", ops.Load())
}

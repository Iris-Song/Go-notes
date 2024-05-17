package main

import (
	"fmt"
	"sync"
)

// For more complex state we can use a mutex to safely access data across multiple goroutines.
// Container holds a map of counters; since we want to update it concurrently from multiple goroutines,
// we add a Mutex to synchronize access.
// Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Lock the mutex before accessing counters; unlock it at the end of the function using a defer statement.
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// Note that the zero value of a mutex is usable as-is, so no initialization is required here.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			c.inc(name)
		}
	}

	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)
	go doIncrement("a", 1000)
	wg.Wait()
	fmt.Println(c.counters)
}

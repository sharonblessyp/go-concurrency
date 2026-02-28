package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	mu  sync.Mutex
	val int
}

/*
mutex maintains data states and ensures only one goroutine can access the critical section at a time.
dont passs mutexes by value, pass by reference (pointer) to ensure all goroutines share the same mutex instance.
*/

func (c *safeCounter) Inc() {
	// Lock so only one goroutine at a time can access the 'value'
	c.mu.Lock()

	// 'defer' ensures the lock is released even if the function panics
	defer c.mu.Unlock()

	c.val++
}

func main() {
	/*
		If you copy a Mutex, you aren't necessarily "racing" for the same memory address
		(because each copy has its own memory); instead, you are breaking the logic of the lock entirely.
		-race flag doesn't detect this
	*/
	counter := &safeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter Value:", counter.val)
}

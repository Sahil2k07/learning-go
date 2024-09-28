package basics3

import (
	"fmt"
	"sync"
)

// Struct with a Read Write Mutex and a data
type SafeData struct {
	mu   sync.RWMutex
	data int
}

// Struct with a mutex and a count
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// The mutex prevents all the write operations but read operations are allowed
func (c *SafeData) read() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data
}

// A function for the mutex to increase its count value
// It prevents the Read as well as Write operations from other go routines while its being modified
func (c *SafeCounter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func mutex() {
	data := SafeData{data: 2}

	fmt.Println("The value of data is:", data.read())

	// Used for having a count of number of Go-Routines actively running
	var wg sync.WaitGroup

	// Initialized struct
	counter := SafeCounter{}

	// Start incrementing the count
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			// Defer keyword puts this function in a stack memory,
			// So this function gets executed right after the anonymous function has completed doing it's things
			// And it executes anyways even if the function panics or gives erronous values
			defer wg.Done()

			counter.increment()
		}()
	}

	// Wait till the wg value is 0 i.e. all the Go-Routines has finished executing
	wg.Wait()

	fmt.Println("The count is:", counter.count)
}

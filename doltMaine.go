// To wait for multiple goroutines to finish, we can
// use a *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()

	}
	wg.Wait()
}

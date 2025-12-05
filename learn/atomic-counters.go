package learn

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounters() {
	var ops atomic.Int64

	var wg sync.WaitGroup

	for range 50 {
		wg.Go(func() {
			for range 1000 {
				ops.Add(1)
			}
		})
	}
	wg.Wait()

	fmt.Println("ops:", ops.Load())
}

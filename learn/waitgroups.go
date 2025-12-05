package learn

import (
	"fmt"
	"sync"
	"time"
)

func waitworkers(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup

	for j := 1; j <= 5; j++ {
		wg.Go(func() {
			waitworkers(j)
		})
	}
	wg.Wait()
}

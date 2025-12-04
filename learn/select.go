package learn

import (
	"fmt"
	"time"
)

func Select() {
	c1 := make(chan bool)
	c2 := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- true
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- false
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("boolean status1:", msg1)
		case msg2 := <-c2:
			fmt.Println("boolean status2:", msg2)
		}
	}
}

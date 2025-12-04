package learn

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working")
	str := "......."
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("awaiting your best output")
			time.Sleep(time.Second)
		}
		fmt.Println(str)
		time.Sleep(time.Second)
	}
	fmt.Println("done")

	done <- true
}

func ChannelSynchro() {
	done := make(chan bool)
	go worker(done)

	<-done
}

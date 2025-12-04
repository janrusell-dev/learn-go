package learn

import (
	"fmt"
	"time"
)

func str(from string) {
	for i := range 5 {
		fmt.Println(from, ":", i)
	}
}

func GoRoutine() {
	str("direct")

	go str("indirect")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

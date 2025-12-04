package learn

import "fmt"

func RangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "69"
	queue <- "420"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

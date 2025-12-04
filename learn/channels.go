package learn

import "fmt"

func Channel() {
	messages := make(chan string)

	go func() {
		messages <- "hehe"
		close(messages)
	}()

	for msg := range messages {
		for _, ch := range msg {
			fmt.Printf("%c\n", ch)
		}
	}
}

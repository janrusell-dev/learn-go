package learn

import (
	"fmt"
	"os"
)

func CLArguments() {
	argsWithProg := os.Args
	argsWithoutProg := argsWithProg[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

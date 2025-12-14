package learn

import (
	"fmt"
	"os"
	"strings"
)

func EnvironmentVariables() {
	os.Setenv("IP_ADDRESS", "192.168.69.69")
	fmt.Println("IP ADDRESS:", os.Getenv("IP_ADDRESS"))
	fmt.Println("LOCALHOST:", os.Getenv("LOCALHOST"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}

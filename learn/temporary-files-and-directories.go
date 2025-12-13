package learn

import (
	"fmt"
	"os"
)

func TempoFilesAndDirectories() {
	f, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write()
}

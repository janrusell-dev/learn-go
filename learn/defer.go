package learn

import (
	"fmt"
	"os"
	"path/filepath"
)

func Defer() {
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func writeFile(f *os.File) {
	fmt.Println("writing..")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing....")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func createFile(path string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

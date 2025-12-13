package learn

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func WritingFiles() {
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat69")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	path2 := filepath.Join(os.TempDir(), "dat420")
	f, err := os.Create(path2)
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

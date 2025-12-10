package learn

import (
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

	d2
}

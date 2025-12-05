package learn

import (
	"os"
	"path/filepath"
)

func Panic() {
	panic("a problem")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
}

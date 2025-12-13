package learn

import (
	"fmt"
	"path/filepath"
	"strings"
)

func FilePaths() {
	p := filepath.Join("dir1", "dir2", "dir3")
	fmt.Println(p)

	fmt.Println(filepath.Join("dir1//", "hehe"))
	fmt.Println(filepath.Join("dir1/../dir1", "asdsad"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}

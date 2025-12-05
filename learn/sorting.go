package learn

import (
	"fmt"
	"slices"
)

func Sorting() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("sorted letters:", strs)

	ints := []int{1, 2, 69, 420, 123}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	s := slices.IsSorted(strs)
	fmt.Println("Sorted: ", s)
}

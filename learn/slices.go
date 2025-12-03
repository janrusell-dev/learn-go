package learn

import (
	"fmt"
	"slices"
)

func Slices() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "c"
	s[1] = "d"
	s[2] = "e"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "z")
	s = append(s, "x", "y")
	fmt.Println("appended letters:", s)

	c := make([]string, len(s))
	copy(c, s)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}

	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i", "j", "k", "l"}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	}

	twoD := make([][]int, 3)
	for i := range twoD {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

package learn

import "fmt"

func Function() {
	result := plus(69, 420)
	fmt.Println("69 + 420 =", result)

	result2 := plusPlus(10, 20, 30)
	fmt.Println("10 + 20 + 30 =", result2)
}

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

package learn

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("after zeroVal:", i)

	zeroPtr(&i)
	fmt.Println("after zeroPtr:", i)

	fmt.Println("pointer:", &i)
}

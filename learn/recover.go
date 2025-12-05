package learn

import "fmt"

func mayPanic() {
	panic("asdasd")
}
func Recover() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()

	fmt.Println("After mayPanic()")
}

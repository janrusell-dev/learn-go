package learn

import "fmt"

type BMI struct {
	weight, height float64
}

func (b *BMI) calculate() float64 {
	return b.weight / (b.height * b.height)
}

func Methods() {
	bmi := BMI{weight: 72, height: 1.83}
	fmt.Println("your bmi is:", bmi.calculate())

	s := &bmi
	fmt.Println("your bmi is:", s.calculate())

}

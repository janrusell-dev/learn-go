package learn

import (
	"fmt"
)

type GetBMI interface {
	calculate() float64
	getCategory() string
}

type bmi struct {
	weight, height float64
}

func (b bmi) calculate() float64 {
	return b.weight / (b.height * b.height)
}

func (b bmi) getCategory() string {
	bmiVal := b.calculate()

	switch {
	case bmiVal < 18.5:
		return "Underweight"
	case bmiVal >= 18.5 && bmiVal <= 24.9:
		return "Normal Weight"
	case bmiVal >= 25 && bmiVal <= 29.9:
		return "Overweight"
	case bmiVal >= 30 && bmiVal <= 34.9:
		return "Obesity class I"
	case bmiVal >= 35 && bmiVal <= 39.9:
		return "Obesity class II"
	default:
		return "Obesity class III"
	}
}

func printWeight(g GetBMI) {
	fmt.Println("your bmi is:", g.calculate())
	fmt.Println("your bmi category was: ", g.getCategory())
}

func Interfaces() {
	p := bmi{weight: 72, height: 1.83}
	printWeight(p)
}

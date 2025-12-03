package learn

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *person {
	p := person{name: name}
	p.age = 23
	return &p
}

func Structs() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Rusell", age: 23})

	fmt.Println(person{name: "Jan"})

	fmt.Println(&person{name: "Jonah", age: 23})

	s := person{name: "sean", age: 78}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 24
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		name:   "Buddy",
		isGood: true,
	}
	fmt.Println(dog.name, dog.isGood)
}

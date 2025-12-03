package learn

import "fmt"

func Range() {
	nums := []int{7, 150, 10, 200}
	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println("sum of a numbers is:", sum)
	}

	for i, num := range nums {
		if num == 7 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key value:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

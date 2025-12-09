package learn

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func JSON() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.4)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("go")
	fmt.Println(string(strB))

	slcD := []string{"asdsd", "asdasdas", "gsdfdsf"}
	slcB, _ := json.Marshal(slcD)
	fmt.Printf(string(slcB))

	mapD := map[string]int{"asdasd": 69, "fgfdgfd": 420}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"asdasd", "asdsad"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	byt := []byte(`{"num": 6.12, "strs": ["a", "b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)

}

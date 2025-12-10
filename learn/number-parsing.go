package learn

import (
	"fmt"
	"strconv"
)

func NumberParsing() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("575", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("12321")
	fmt.Println(k)

	_, e := strconv.Atoi("Wait")
	fmt.Println(e)

}

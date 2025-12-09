package learn

import (
	"fmt"
	"os"
	"time"
)

type point struct {
	x, y int
}

func StringFormatting() {
	for i := 0; i <= 1000; i++ {

		p := point{69, 420}
		fmt.Printf("struct 1: %v\n", p)

		fmt.Printf("struct 2: %+v\n", p)

		fmt.Printf("struct 3: %#v\n", p)

		fmt.Printf("type: %T\n", p)

		fmt.Printf("bool: %t\n", true)

		fmt.Printf("int: %d\n", 69)

		fmt.Printf("bin: %b\n", 14)

		fmt.Printf("char: %c\n", 67)

		fmt.Printf("hex: %x\n", 456)

		fmt.Printf("float 1: %f\n", 12.24)

		fmt.Printf("float2: %e\n", 123400000.0)
		fmt.Printf("float3: %E\n", 123400000.0)

		fmt.Printf("str1 : %s\n", "\"string\"")
		fmt.Printf("str2: %q\n", "\"string\"")
		fmt.Printf("str2: %x\n", "asdsd")

		fmt.Printf("pointer: %p\n", &p)

		fmt.Printf("width 2: |%6d|%6d|\n", 12, 562)

		fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

		fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

		s := fmt.Sprintf("sprintf: a %s", "sadsda")
		fmt.Println(s)

		fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

		fmt.Printf("Youre on line %v", i)
		time.Sleep(time.Second)
	}
}

package src

import "fmt"

var str1 string = "Hello"

const (
	A int = iota
	B int = iota
)

func test() {
	var x1 int = 1
	if x1 > 5 {
		fmt.Print("H")
	}

	for i := 0; i < 10; i++ {
		fmt.Print("H")
	}

	switch {
	case x1 == 1:
		fmt.Print("H")

	default:
		fmt.Print("H")

	}

}

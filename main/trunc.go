package main

import (
	"fmt"
	"strings"
)

func main() {
	NumFloat := ""
	fmt.Println("Enter a floating point number. ")
	_, _ = fmt.Scan(&NumFloat)
	FloatingPart := strings.Split(NumFloat, ".")[1]
	NumFloat = "0." + FloatingPart
	fmt.Println(NumFloat)
}

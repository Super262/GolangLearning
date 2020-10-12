package main

import "fmt"

func main() {
	fmt.Println("Enter a floating point number. ")
	var numFloat float64 = 0.0
	var numInt int64 = 0
	fmt.Scan(&numFloat)
	numInt = int64(numFloat)
	fmt.Print(numFloat - float64(numInt))
}

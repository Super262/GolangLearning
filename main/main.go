package main

import (
	"fmt"
)

func main() {
	var x = [...]int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Print(i)
		fmt.Print(" ")
		fmt.Print(v)
		fmt.Println()
	}
	slice1 := x[0:1]
	fmt.Printf("Capacity of slice1: %d, length of slice1: %d.\n", cap(slice1), len(slice1))
	slice2 := make([]int, 10, 15)
	slice2 = append(slice2, 100)
	fmt.Printf("Capacity of slice1: %d, length of slice1: %d.\n", cap(slice2), len(slice2))
	fmt.Println(slice2)

	idMap := make(map[string]int)
	idMap["Joe"] = 456
	idMap["Green"] = 123
	fmt.Printf("Joe's ID is %d.\n", idMap["Joe"])
	fmt.Printf("Green's ID is %d.\n", idMap["Green"])
	delete(idMap, "Joe")
	id, p := idMap["Joe"]
	fmt.Println(p)
	fmt.Printf("Joe's ID is %d.\n", id)
	idMap["Joe"] = 456
	idMap["Green"] = 123
	idMap["Peter"] = 789
	idMap["Ann"] = 012
	idMap["Jane"] = 345
	idMap["Lisa"] = 678
	for key, value := range idMap {
		fmt.Println(key, value)
	}
}

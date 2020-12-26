package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func less(i int, j int) bool {
	if i < j {
		return true
	} else {
		return false
	}
}

func main() {
	slice := make([]int, 0, 3)
	var input = "0"
	for {
		fmt.Println("Enter an integer. Enter X to quit.")
		_, _ = fmt.Scan(&input)
		if strings.Compare(input, "X") == 0 {
			break
		}
		val, _ := strconv.Atoi(input)
		slice = append(slice, val)
		sort.Ints(slice)
		fmt.Println(slice)
	}
}

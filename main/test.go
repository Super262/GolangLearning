package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	sl := make([]int, 0, 3)
	for {
		fmt.Print("Enter the new number(Enter X to exit) :: ")
		inp, _ := reader.ReadString('\n')
		inp = strings.TrimRight(inp, "\r\n")
		inp = strings.ToUpper(inp)
		if inp == "X" {
			break
		} else {
			inp1, _ := strconv.Atoi(inp)
			sl = append(sl, inp1)
			sort.Ints(sl)
			fmt.Printf("%d\n", sl)
		}
	}

}

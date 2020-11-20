package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string. ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	StrEntered := scanner.Text()
	RuneEntered := []rune(StrEntered)
	LenRune := len(RuneEntered)

	if strings.Compare(string(RuneEntered[0]), "i") != 0 && strings.Compare(string(RuneEntered[0]), "I") != 0 {
		fmt.Println("Not Found!")
		return
	}

	if strings.Compare(string(RuneEntered[LenRune-1]), "n") != 0 && strings.Compare(string(RuneEntered[LenRune-1]), "N") != 0 {
		fmt.Println("Not Found!")
		return
	}

	if strings.Contains(StrEntered, "a") || strings.Contains(StrEntered, "A") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

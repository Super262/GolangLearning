package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Filename: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()
	file, err1 := os.Open(filename)
	if err1 == nil {
		reader := bufio.NewReader(file)
		nameSlice := make([]Name, 0, 100)
		for {
			line, err := reader.ReadString('\n')
			if strings.Compare(line, "") == 0 || (err != nil && err != io.EOF) {
				break
			} else {
				content := strings.Split(line, " ")
				name := Name{fname: strings.TrimSpace(content[0]), lname: strings.TrimSpace(content[1])}
				nameSlice = append(nameSlice, name)
			}
		}
		for _, name := range nameSlice {
			fmt.Println("First name: " + name.fname + ", Last name: " + name.lname)
		}
	}
}

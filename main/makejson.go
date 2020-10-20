package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	DataMap := make(map[string]string)
	fmt.Print("Name: ")
	Scanner.Scan()
	name := Scanner.Text()
	fmt.Print("Address: ")
	Scanner.Scan()
	address := Scanner.Text()
	DataMap["address"] = address
	DataMap["name"] = name
	JsonObj, err := json.Marshal(DataMap)
	if err == nil {
		fmt.Println(JsonObj)
	}
}

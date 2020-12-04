package main

import (
	"fmt"
	"strconv"
	"time"
)

func printAll(counter *int, name string, idle int) {
	for i := 0; i < 10; i++ {
		*counter = *counter + 1
		fmt.Println("Goroutine " + name + ": " + strconv.Itoa(*counter))
		time.Sleep(time.Duration(idle) * time.Millisecond)
	}
}

func main() {
	var counter int
	counter = 0
	go printAll(&counter, "1", 125)
	go printAll(&counter, "2", 250)
	time.Sleep(5000 * time.Millisecond)
}

// These two threads are executing the similar code from the start respectively.
// Although the final result is proper, a race condition occurs during the whole process.

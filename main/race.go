package main

import (
	"fmt"
	"time"
)

func setCounter(counter *int, idle int) {
	for i := 0; i < 10; i++ {
		value := *counter
		time.Sleep(time.Duration(idle) * time.Millisecond)
		value++
		*counter = value
	}
}

func main() {
	var counter int
	counter = 0
	go setCounter(&counter, 125)
	go setCounter(&counter, 250)
	time.Sleep(5000 * time.Millisecond)
	fmt.Printf("Final Value of Counter: %d\n", counter)
}

// Explanations:
// These two threads are executing the same code from the start, respectively.
// The final result, which should be 20, is 10. The final result is wrong because a race condition occurs during the whole process.
// Every query would get the old value of counter during the whole process because these two threads are performing write operations without any limits.

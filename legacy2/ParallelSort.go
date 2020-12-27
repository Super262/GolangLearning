package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortOnePart(number int, slice []int, wg *sync.WaitGroup) {
	fmt.Printf("goroutine %d's subarray: %v\n", number, slice)
	sort.Ints(slice)
	wg.Done()
}

func main() {
	slicesNum := 4
	var arrayLength int
	fmt.Printf("Enter an integer which must be larger than %d to set the length of the array: ", slicesNum-1)
	_, _ = fmt.Scan(&arrayLength)
	if arrayLength < slicesNum {
		fmt.Println("The length is too small! ")
	} else {
		var wg sync.WaitGroup
		inputArray := make([]int, arrayLength, arrayLength)
		sliceLength := arrayLength / slicesNum

		// initialization
		fmt.Printf("Enter %d integers as elements of the array: ", arrayLength)
		for i := 0; i < arrayLength; i++ {
			_, _ = fmt.Scan(&inputArray[i])
		}

		// start goroutine
		wg.Add(4)
		for i := 0; i < slicesNum-1; i++ {
			go sortOnePart(i, inputArray[i*sliceLength:(i+1)*sliceLength], &wg)
		}
		go sortOnePart(slicesNum-1, inputArray[(slicesNum-1)*sliceLength:arrayLength], &wg)

		// Merge and Sort
		wg.Wait()
		sort.Ints(inputArray)

		fmt.Print("All elements of the sorted array: ")
		for i := 0; i < arrayLength; i++ {
			fmt.Printf("%d ", inputArray[i])
		}
		fmt.Println()
	}
}

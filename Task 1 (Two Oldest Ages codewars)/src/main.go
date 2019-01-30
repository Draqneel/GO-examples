package main

import (
	"fmt"
	"math"
)

func main() {
	var count int
	var value int
	slice := make([]int, 0, 10)
	fmt.Println("Add count of elements")
	fmt.Scan(&count)

	// add elements to slice
	for i := 0; i < count; i++ {
		fmt.Println("Add ", (i + 1), " element")
		fmt.Scan(&value)
		slice = append(slice, value)
	}

	maxOne, maxTwo := findTwoOldestValue(slice)
	fmt.Println("Max 1:", maxOne, "\n", "Max 2:", maxTwo)
}

func findTwoOldestValue(slice []int) (int, int) {
	maxOne := slice[0]
	maxTwo := math.MinInt8
	for i, element := range slice {
		// in default Max One element equals slice[0]
		// we don't need to check it
		if i == 0 {
			continue
		}

		if element > maxOne {
			maxTwo = maxOne
			maxOne = element
			continue
		}

		if element > maxTwo {
			maxTwo = element
		}
	}

	return maxOne, maxTwo
}

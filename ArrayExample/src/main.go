package main

import "fmt"
import "math"

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

	max := slice[0]
	min := math.MaxInt8
	mid := math.MaxInt8

	// find min and max elements
	findMaxMinInSlice(slice, &max, &min)
	// find mid element
	findMidValueInSlice(slice, max, min, &mid)

	// output in console
	fmt.Println(" Min value:", min, "\n", "Mid value", mid, "\n", "Max value:", max)
	printIntSlice(slice)
}

func printIntSlice(slice []int) {
	fmt.Println("all elements of slice:")
	for i, element := range slice {
		fmt.Println((i + 1), "Element:", element)
	}
}

func findMaxMinInSlice(slice []int, max *int, min *int) {
	for _, element := range slice {
		if element > *max {
			*max = element
		}

		if element < *min {
			*min = element
		}
	}
}

func findMidValueInSlice(slice []int, max int, min int, mid *int) {
	delta := (max + min) / 2

	for _, element := range slice {
		if (*mid <= delta) && (element >= delta) {
			if (delta - *mid) > (element - delta) {
				*mid = element
			}
		} else if (*mid <= delta) && (element <= delta) {
			if element > *mid {
				*mid = element
			}
		}

		if (*mid >= delta) && (element <= delta) {
			if (delta - element) < (*mid - delta) {
				*mid = element
			}
		} else if (*mid >= delta) && (element >= delta) {
			if element < *mid {
				*mid = element
			}
		}
	}
}

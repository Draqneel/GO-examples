package main

import "fmt"

// learning how to create:
// functions (simple or recursion
// slices with needed size and capacity
// get length and capacity of slice/array
func main() {
	// slice == ArrayList in GO
	ArrayList := make([]int, 0, 10)
	i := 0
	for i < 5 {
		ArrayList = append(ArrayList, i)
		i++
	}
	fmt.Println(ArrayList, "count of elements in slice", len(ArrayList), "capacity of slice", cap(ArrayList))
	fmt.Println("factorial 5 is", recursion(5))
	fmt.Println(expToTwo(3, 10))
}

// factorial
func recursion(value uint) uint {
	if value == 0 || value == 1 {
		return 1
	}
	return value * recursion(value-1)
}

// value * 2^exp
func expToTwo(value uint, exp uint) uint {
	return value << exp
}

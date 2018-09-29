package main

import "fmt"

// learning how to create:
// functions (simple or recursion
// slices with needed size and capacity
// get length and capacity of slice/array
func main() {
	// WORK WITH SLICES
	// slice == ArrayList in GO () => slice is reference type
	// in variable "ArrayList" created slice ([]int) with size (0) and capacity (10)
	ArrayList := make([]int, 0, 10)
	i := 0
	for i < 5 {
		// we can append not only elements of current data type
		// and yet other slice (all elements from 2nd slice will be added in 1st)
		// and parts of slices
		ArrayList = append(ArrayList, 5)
		i++
	}
	// create slice two with size equals slice one
	ArrayListTwo := make([]int, len(ArrayList))
	// copy elements from "ArrayLIst" to "ArrayListTwo"
	copy(ArrayListTwo, ArrayList)
	fmt.Println(ArrayList, "count of elements in slice", len(ArrayList), "capacity of slice", cap(ArrayList))
	fmt.Println("If we try to copy elements from slice one to slice two we get result:",ArrayListTwo)
	// output from [0:3) - 0, 1, 2
	fmt.Println("Output part of slice",ArrayListTwo	[:3])
	// output from [2:4) - 2, 3
	fmt.Println("Output part of slice",ArrayListTwo	[2:4])

	// WORK WITH FUNCTIONS
	fmt.Println("factorial 5 is", recursion(5))
	fmt.Println(expToTwo(3, 10))

	//WORK WITH MAPS
	// incorrect map creating <code> var HashMap map[string]int </code>(ref == nil, we cant added element)
	// correct map creating:
	HashMap := map[string]int{} // or <code> var HashMap map[string]int = map[string]int{} </code>
	HashMap["birthDay"] = 14
	HashMap["birthMonth"] = 4
	HashMap["birthYear"] = 1999
	// empty map = nil (nil == null)
	fmt.Println("Uninitialized map", HashMap)
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

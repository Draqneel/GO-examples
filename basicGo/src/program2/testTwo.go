package main

import "fmt"

var stuff = "not ready"

func main() {
	// testing functions
	// defer called when func that include it ended
	defer fmt.Println("PROGRAM ENDED")
	arguments := []int{100, -20, 30}
	fmt.Println(sum(arguments...), "sum of all elements of slice")
	fmt.Println("The stuff is", stuff)
	// testing * and &
	// create variable of int data type
	variable := 1
	// result of &NAME_OF_VAR = *int data type (pointer)
	inc(&variable)
	fmt.Println(variable)
	// test of panic
	panicTest()
}

func init() {
	// starting faster than "main" (analogue of "static" methods in Java)
	stuff = "ready"
	fmt.Println("PROGRAM STARTED")
}

func sum(argument ...int) (sum int) {
	// "_" - index in slice, "value" - value of slice's element
	for _, value := range argument {
		sum += value
	}
	return
}

func inc(x *int) {
	*x = *x + 1
}

func panicTest() {
	// error processing
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC UNDER CONTROL!!!")
		}
	}()
	fmt.Println("Start of panicTest()")
	// create error
	panic("At the disco")
	fmt.Println("HAhaHAhAhAH")
}

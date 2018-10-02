package main

import "fmt"

func main() {
	fmt.Println("Simple calc application")
	fmt.Println("1) Addition")
	fmt.Println("2) Subtraction")
	fmt.Println("3) Division")
	fmt.Println("4) Multiplication")
	var operation int
	fmt.Scan(&operation)
	var valueOne int
	var valueTwo int
	fmt.Println("Add value 1")
	fmt.Scan(&valueOne)
	fmt.Println("Add value 2")
	fmt.Scan(&valueTwo)

	switch operation {
	case 1:
		{
			fmt.Println("Result:", valueOne+valueTwo)
		}
	case 2:
		{
			fmt.Println("Result:", valueOne-valueTwo)
		}
	case 3:
		{
			fmt.Println("Result:", valueOne/valueTwo)
		}
	case 4:
		{
			fmt.Println("Result:", valueOne*valueTwo)
		}
	default:
		{
			fmt.Println("Incorrect operation")
		}
	}

}

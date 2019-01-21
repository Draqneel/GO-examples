package main

import "fmt"

/*
 Specification: http://acm.timus.ru/problem.aspx?space=1&num=2095&locale=en
*/
const maxvalue = 1000000000

func main() {
	var (
		fistBoard   int
		secondBoard int
		checkDay    int = 0
		result      int = 0
	)
	fmt.Println("Add 1st value")
	fmt.Scan(&fistBoard)
	fmt.Println("Add 2nd value")
	fmt.Scan(&secondBoard)
	if (fistBoard < 1 || fistBoard > maxvalue) && (secondBoard < 1 || secondBoard > maxvalue) && (fistBoard > secondBoard) {
		fmt.Println("Incorrect input")
		return
	}

	for checkDay < maxvalue {
		if checkDay > secondBoard {
			break
		}
		checkDay++
		if (checkDay >= fistBoard) && (checkDay <= secondBoard) {
			result++
		}
		checkDay += 2
		if (checkDay >= fistBoard) && (checkDay <= secondBoard) {
			result++
		}
		checkDay += 4
		if (checkDay >= fistBoard) && (checkDay <= secondBoard) {
			result++
		}
		checkDay += 3
	}

	fmt.Println("result", result)
}

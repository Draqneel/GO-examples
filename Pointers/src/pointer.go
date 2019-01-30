package main

import "fmt"

func main() {
	str := "Hello, world!"
	for i, value := range str {
		fmt.Printf("Character %c , position %d \n", value, i)
	}
}

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount - specification : https://tour.golang.org/moretypes/23
func WordCount(s string) map[string]int {
	mapp := make(map[string]int)
	str := strings.Split(s, " ")
	for _, value := range str {
		mapp[value]++
	}
	return mapp
}

func main() {
	wc.Test(WordCount)
}

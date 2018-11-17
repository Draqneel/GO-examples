pachage main

import "fmt"

type Printer struct {
	message string
}

func (pr *Printer) print(msg string) {
	fmt.Println(msg)
}
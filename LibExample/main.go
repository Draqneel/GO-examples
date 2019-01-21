package main

import (
	"fmt"
	"sync"
)

// test public struct
type MyObj struct {
	dataMutex *sync.Mutex
	data      []int
	count     int
}

func (obj *MyObj) addData(value int) {
	obj.dataMutex.Lock()
	defer obj.dataMutex.Unlock()
	obj.data[obj.count] = value
	obj.count++
}

func main() {
	obj := MyObj{
		data:      make([]int, 5),
		dataMutex: new(sync.Mutex),
		count:     0,
	}
	obj.addData(10)
	obj.addData(50)
	obj.addData(100)
	fmt.Println(obj.data)
	//

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		go Print("Piece", wg)
	}
	wg.Wait()
}

// test public method
func Print(word string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	fmt.Println(word)
}

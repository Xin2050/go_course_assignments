package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0

// this code will given different result because data race
// use go build -race check it

// add a lock  to fix this problem
var lock sync.Mutex

func main() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}
	Wait.Wait()
	fmt.Println("Final Counter:", Counter)
}

func Routine(id int) {

	lock.Lock()
	for count := 0; count < 2; count++ {
		value := Counter
		time.Sleep(time.Nanosecond)
		value++
		Counter = value
	}
	lock.Unlock()
	Wait.Done()
}

package main

import (
	"github.com/Xin2050/go_course_assignments/s4/test/client"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	wg := sync.WaitGroup{}
	for {
		wg.Add(1)
		go func() {
			client.Client()
			wg.Done()
		}()
		time.Sleep(time.Millisecond * 20)
	}
	wg.Wait()
}

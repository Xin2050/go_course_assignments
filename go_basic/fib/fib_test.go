package fib

import (
	"fmt"
	"testing"
	"time"
)

var dict = make(map[int]int64, 40)

func fibWithCache(n int) int64 {

	if n <= 2 {
		return 1
	}
	if _, ok := dict[n]; !ok {
		dict[n] = int64(fibWithCache(n-1) + fibWithCache(n-2))
	}
	return dict[n]
}
func fib(n int) int64 {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibWithRoutine(n int) int64 {
	if n <= 2 {
		return 1
	}
	//if n < 10 {
	//	return fib(n-1)+fib(n-2)
	//}
	var c = make(chan int64)

	go func() {
		c <- fibWithCache(n - 1)
	}()
	go func() {
		c <- fibWithCache(n - 2)
	}()
	return <-c + <-c
}

func TestFibTime(t *testing.T) {
	testNumb := 150
	startTime := time.Now()
	fmt.Println(fibWithCache(testNumb))
	fmt.Printf("use cache time :%s\n", time.Since(startTime))
	// don't try to run over 45
	//startTime = time.Now()
	//fmt.Println(fib(testNumb))
	//fmt.Printf("no cache time :%s\n", time.Since(startTime))
	startTime = time.Now()
	fmt.Println(fibWithRoutine(testNumb))

	fmt.Printf("go routine with cache time :%s\n", time.Since(startTime))
}

func TestFib(t *testing.T) {

	if fib(1) != 1 {
		t.Fatal("failed to fib(1) = 1")
	}
	if fib(2) != 1 {
		t.Fatal("failed to fib(2) = 1")
	}
	if fib(3) != 2 {
		t.Fatal("failed to fib(3) = 2")
	}
	if fib(39) != 63245986 {
		t.Fatal("failed to fib(39) = 63245986")
	}
}

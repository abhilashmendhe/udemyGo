package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 40; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Millisecond * 3)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 40; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Millisecond * 2)
	}
	wg.Done()
}

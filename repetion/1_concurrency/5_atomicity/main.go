package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64
var mutex sync.Mutex

func main() {

	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
	fmt.Println("Final counter : ", counter)

}

func increment(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Coutner:", atomic.LoadInt64(&counter))
	}
	wg.Done()
}

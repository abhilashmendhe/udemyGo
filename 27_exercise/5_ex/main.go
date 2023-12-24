package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var increment int64
	var wg sync.WaitGroup
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&increment, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(increment)
}

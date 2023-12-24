package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOOS:", runtime.NumGoroutine())

	var count int64 = 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println("Count:", atomic.LoadInt64(&count))
			wg.Done()
		}()
		fmt.Println("GOOS:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("count:", count)
}

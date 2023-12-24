package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOOS:", runtime.NumGoroutine())

	count := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := count
			// time.Sleep(time.Second * 1)
			runtime.Gosched()
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GOOS:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("count:", count)
}

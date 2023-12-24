package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	increment := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := increment
			runtime.Gosched()
			v++
			increment = v
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(increment)
}

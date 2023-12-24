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
	var m sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {

			m.Lock()
			v := increment
			runtime.Gosched()
			v++
			increment = v
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(increment)
}

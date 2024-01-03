package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
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
		mutex.Lock()
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

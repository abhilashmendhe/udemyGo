package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {

	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
	fmt.Println("Final counter : ", counter)

}

func increment(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

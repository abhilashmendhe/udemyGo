package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}
func doWork() {
	for i := 0; i < 10000000000; i++ {

	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	end := time.Since(start)
	fmt.Println(end)
}

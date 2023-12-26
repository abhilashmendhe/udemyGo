package main

import "fmt"

func main() {
	// cs := make(chan<- int) // this is only send only channel,
	// 						// remove arrow to make it bidirectional channel
	cs := make(chan int)
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
}

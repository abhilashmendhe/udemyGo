package main

import "fmt"

func main() {

	c := make(chan int)
	// send
	go foo(c)
	// receive
	// bar(c)
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Exit!!")
}

// send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- 100
	}
	close(c) // if using range to read channel then definitely use close(chan) func.
}

// receive
// func bar(c <-chan int) {
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(i, "-", <-c)
// 	}
// }

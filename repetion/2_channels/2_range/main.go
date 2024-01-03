package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // when this is called, you can no longer put values on
		// the channel
	}()

	for n := range c {
		fmt.Println(n)
	}
}

package main

import (
	"fmt"
)

func main() {

	// single channel
	/*	c := make(chan int)
		go func() {
			c <- 42
		}()

		fmt.Println(<-c)
	*/
	// buffer channel
	c := make(chan int, 4)
	c <- 41
	c <- 42
	c <- 43
	c <- 44
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

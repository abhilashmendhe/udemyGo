package main

import "fmt"

func main() {
	// pull the values off the channel using for range loop
	c := gen()
	receive(c)
	fmt.Println("exit!!")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 15; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

package main

import "fmt"

/*
write a program
- launches 10 goroutines
  - each goroutine adds 10 numbers to channel

- pull the numbers off the channel and print them
*/
func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {

		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}
	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}
}

package main

import "fmt"

func main() {

	c := make(chan int)
	cr := make(<-chan int) // recieve
	cs := make(chan<- int) // send
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
	// c := make(chan int, 2)
	// c <- 1
	// c <- 2
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println("---------")
	// fmt.Printf("%T\n", c)

}

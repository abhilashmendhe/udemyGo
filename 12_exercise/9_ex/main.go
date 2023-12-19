package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("Outer loop: %v\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("Inner loop\n")
			fmt.Printf("%v ", j)
		}
	}
}

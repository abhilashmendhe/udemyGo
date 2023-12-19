package main

import "fmt"

func main() {

	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("xi - %#v\n", xi)

	// [inclusive:exclusive]
	fmt.Printf("xi - %#v\n", xi[4:7])

	// [:exclusive]
	fmt.Printf("xi - %#v\n", xi[:5])

	// [inclusive:]
	fmt.Printf("xi - %#v\n", xi[5:])
}

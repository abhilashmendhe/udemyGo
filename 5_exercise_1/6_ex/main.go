package main

import "fmt"

func main() {
	a, b, c := 747, 911, 90210

	fmt.Printf("%d : %d : %#x\n", a, a, a)
	fmt.Printf("%d : %b : %#x\n", b, b, b)
	fmt.Printf("%d : %b : %#x\n", c, c, c)
}

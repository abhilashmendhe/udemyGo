package main

import "fmt"

func main() {

	// array literal
	a := [3]int{42, 21, 55}
	fmt.Println(a)

	b := [...]string{"Hello", "Gopher"}
	fmt.Printf("%T\n", b)

	var c [2]int
	c[0] = 7
	c[1] = 10
	fmt.Printf("%T\n", c)

}

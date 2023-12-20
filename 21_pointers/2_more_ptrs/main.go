package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%v\t%T\n", &x, &x)
	s := "abhi"
	fmt.Printf("%v\t%T\n", &s, &s)
}

package main

import "fmt"

func main() {
	s, e, f := "Happiness", 32, 43.2
	fmt.Printf("%v is of type %T\n", s, s)
	fmt.Printf("%v is of type %T\n", e, e)
	fmt.Printf("%v is of type %T\n", f, f)
}

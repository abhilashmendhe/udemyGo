package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y := bar()
	y2 := y()
	fmt.Println(y2)
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}

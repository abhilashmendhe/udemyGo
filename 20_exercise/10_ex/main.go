package main

import "fmt"

func main() {
	x := outer()
	fmt.Println(x())
}

func outer() func() int {
	return func() int {
		return 44
	}
}

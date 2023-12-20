package main

import "fmt"

func main() {
	f := foo()
	b1, b2 := bar()
	fmt.Println(f)
	fmt.Println(b1, b2)
}

func foo() int {
	return 42
}
func bar() (int, string) {
	return 100, "hello gophers"
}

package main

import "fmt"

func main() {
	defer foo()
	bar()
}

// func (r reciever) identifier(p params(s)) (r returns(s)){ code }

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}

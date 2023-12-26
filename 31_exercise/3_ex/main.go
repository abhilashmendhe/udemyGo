package main

import (
	"fmt"
)

/*
Create a struct custom error..
*/
type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}
func main() {
	c1 := customErr{
		info: "I need food",
	}
	foo(c1)
}
func foo(e error) {
	fmt.Println("foo ran - ", e, "\n")
}

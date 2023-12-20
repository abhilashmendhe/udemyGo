package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}
type count int

func (b book) String() string {
	return fmt.Sprint("this is the book is ", b.title)
}

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "West with the night",
	}
	var c count = 42
	fmt.Println(c)
	fmt.Println(b)
}

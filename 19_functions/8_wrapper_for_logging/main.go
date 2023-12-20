package main

import (
	"fmt"
	"log"
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
func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}
func main() {
	b := book{
		title: "West with the night",
	}
	var c count = 42
	logInfo(b)
	logInfo(c)
}

package main

import (
	"bytes"
	"fmt"
)

func main() {

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())

	b.WriteString("Gohpers!!")
	fmt.Println(b.String())

	b.Reset()

	b.WriteString("I am alone!!")
	fmt.Println(b.String())

	b.Write([]byte("Happy Happy"))
	fmt.Println(b.String())
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	first string
}

func (p Person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {

	p := Person{
		first: "Jenny",
	}
	p2 := Person{
		first: "abhi",
	}
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	p2.writeOut(f)
	p2.writeOut(&b)
	fmt.Println(b.String())
}

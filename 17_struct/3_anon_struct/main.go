package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
	}{
		first: "abhi",
		last:  "menda",
	}
	fmt.Println(p1)
}

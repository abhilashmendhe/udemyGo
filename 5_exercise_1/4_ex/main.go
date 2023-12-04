package main

import "fmt"

var y int

func main() {
	fmt.Println(y)

	z := 42
	fmt.Println(z)

	a, b := 32, "adsfasdf"
	fmt.Println(a, b)

	var c float32 = 123.123
	fmt.Printf("%v is of this type %T", c, c)

	e, f, _ := 12, 32.1, "hoasd"
	fmt.Println(e, f)
}

/*
var for zero value
short declaration operator
multiple initialization
var when specificity is requireed
blank identifier
*/

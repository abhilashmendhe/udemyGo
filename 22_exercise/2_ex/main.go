package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("%v type is %T\n", *a, a)
	fmt.Printf("%v type is %T\n", *b, b)
	fmt.Printf("%v type is %T\n", *c, c)
	fmt.Printf("%v type is %T\n", *d, d)
	// fmt.Printf("%v type is %T\n", p, p)
}

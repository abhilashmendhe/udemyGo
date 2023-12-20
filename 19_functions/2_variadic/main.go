package main

import "fmt"

func main() {
	s := sum("abhi", 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(s)
	name := []string{"abc", "efg", "dec"}
	names(name...)
}

func names(s ...string) {
	fmt.Println(s)
}
func sum(s string, ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

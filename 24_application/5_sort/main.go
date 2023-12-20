package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 55}
	sort.Ints(xi)
	fmt.Println(xi)
}

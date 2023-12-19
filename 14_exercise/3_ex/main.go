package main

import "fmt"

func main() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(sl)

	fmt.Println(sl[:5])
	fmt.Println(sl[5:])
	fmt.Println(sl[2:7])
	fmt.Println(sl[3:8])
}

package main

import "fmt"

func main() {

	am := make(map[string]int)
	am["abhi"] = 29
	am["john"] = 32
	am["sarah"] = 27
	for k, v := range am {
		fmt.Println(k, "age is", v)
	}
}

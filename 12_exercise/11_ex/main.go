package main

import "fmt"

func main() {
	m := map[string]int{
		"Abhi": 29,
		"Lala": 25,
	}

	for k, v := range m {
		fmt.Println(k, ": ", v)
		fmt.Println(m["Q"])
	}

}

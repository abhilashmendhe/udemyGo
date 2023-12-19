package main

import "fmt"

func main() {

	am := map[string]int{
		"abhi": 29,
		"john": 32,
	}
	fmt.Println(am)
	fmt.Printf("The age of abhi is %v\n", am["abhi"])

	map2 := make(map[string]string)
	map2["abhi"] = "menda"
	map2["john"] = "pertruci"
	fmt.Printf("%#v\n", map2)

	fmt.Println("Len of map: ", len(map2))
}

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Abhi",
		Last:  "Menda",
		Age:   29,
	}
	p2 := person{
		First: "Vasiliki",
		Last:  "Letta",
		Age:   25,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	//make sure the fields in struct
	// start with capital letter
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(string(bs))
}

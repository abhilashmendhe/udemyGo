package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	Create custom error message using fmt.Errorf
*/

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Hahahah", "You die!!"},
	}
	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("There is an error %v", err)
	}
	return bs, nil
}

package main

import (
	"fmt"
	"strings"
)

// kv := make(map[string]int)
var kv = make(map[string]int)

func main() {

	kv["k1"] = 10
	kv["k2"] = 12
	fmt.Println(kv)
	val, is := kv["k3"]
	fmt.Println(val, is)
	delete(kv, "k1")
	fmt.Println(kv)

	s := "GET kv1"
	s_arr := strings.Split(s, " ")
	fmt.Println(s_arr)
}

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
	delete(am, "john")
	fmt.Println("\nDeleted john\n")
	for k, v := range am {
		fmt.Println(k, "age is", v)
	}

	// checkif key value exist
	v, ok := am["john"]
	if ok {
		fmt.Println("Key john exists", v)
	} else {
		fmt.Println("john does not exist")
	}

	// comman ok idiom
	if v, ok := am["abhi"]; ok {
		fmt.Println("abhi exists", v)
	} else {
		fmt.Println("abhi does not exists")
	}
}

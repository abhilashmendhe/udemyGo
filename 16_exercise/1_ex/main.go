package main

import "fmt"

func main() {
	m1 := make(map[string][]string)
	m1["bond_james"] = append(m1["bond_james"], "shaken, not stirred", "martinis", "fast cars")

	arr := []string{"intelligence", "literature", "computer science"}
	m1["moneypenny_jenny"] = append(m1["moneypenny_jenny"], arr...)

	m1["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for k, v := range m1 {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, "->", v2)
		}
	}
}

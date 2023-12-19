package main

func main() {

	m1 := make(map[string][]string)
	m1["bond_james"] = append(m1["bond_james"], "shaken, not stirred", "martinis", "fast cars")

	arr := []string{"intelligence", "literature", "computer science"}
	m1["moneypenny_jenny"] = append(m1["moneypenny_jenny"], arr...)

	m1["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	m1["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./book.txt")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Book exists")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	wordcount := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		for _, w := range words {
			if _, ok := wordcount[w]; !ok {
				wordcount[w] = 1
			} else {
				wordcount[w] += 1
			}
		}
	}

	for k, v := range wordcount {
		fmt.Println(k, ":", v)
	}
}

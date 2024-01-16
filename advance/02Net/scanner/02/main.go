package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}
	fmt.Println(count)
}

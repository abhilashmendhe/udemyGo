package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("-> ")
	// 	text, _ := reader.ReadString('\n')
	// 	fmt.Print(text)
	// }

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("-> ")
	for scanner.Scan() {

		fmt.Println(scanner.Text())
		fmt.Print("-> ")
	}
}

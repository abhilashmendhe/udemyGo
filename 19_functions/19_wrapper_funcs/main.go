package main

import (
	"fmt"
	"log"
	"os"
)

// THIS CODE WON'T RUN ON GO PLAYGROUND
// because it has file operations
// you can read it though

func main() {
	xb, err := readFile("output.txt")
	if err != nil {
		log.Fatalf("error in main in readFile: %s", err)
	}
	fmt.Println(xb) // print int arr which are bytes
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("there was an error in readfile: %s", err)
	}
	return xb, nil
}

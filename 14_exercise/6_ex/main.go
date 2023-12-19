package main

import "fmt"

func main() {
	jb := []string{"abc", "bbc", "cca"}
	jm := []string{"ccc", "cbc", "dca"}

	mulsl := [][]string{jb, jm}
	fmt.Println(mulsl)
}

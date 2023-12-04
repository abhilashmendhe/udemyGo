package futherexplored

import "fmt"

// This will be exported since starting with capital letter
func Fascinating() {
	fmt.Println("Planet Speed -> ", planetSpeed)

	planetRotationSpeed := 1000
	fmt.Println("Planet spinning speed: ", planetRotationSpeed)
}

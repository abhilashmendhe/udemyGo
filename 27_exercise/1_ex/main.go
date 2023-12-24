package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin CPUs:", runtime.NumCPU())
	fmt.Println("Begin goos:", runtime.NumGoroutine())
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("This is first.")
		wg.Done()
	}()

	go func() {
		fmt.Println("This is second.")
		wg.Done()
	}()
	fmt.Println("Mid CPUs:", runtime.NumCPU())
	fmt.Println("Mid goos:", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Final CPUs:", runtime.NumCPU())
	fmt.Println("Final goos:", runtime.NumGoroutine())
	fmt.Println("\nFinished printing 2 goroutines!!!")
}

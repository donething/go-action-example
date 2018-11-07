package main

import (
	"fmt"
	"sync"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for c := 'a'; c < 'a'+26; c++ {
				fmt.Printf("%c ", c)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

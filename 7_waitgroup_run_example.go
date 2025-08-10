package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Old method:")
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		fmt.Println("Goroutine (old method)")
	}()
	wg1.Wait()

	// new method
	fmt.Println("\nNew method (Go 1.25+):")
	var wg2 sync.WaitGroup
	wg2.Go(func() {
		fmt.Println("Goroutine (new method)")
	})
	wg2.Wait()
}

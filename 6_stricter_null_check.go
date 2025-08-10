package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("file_invalid.txt") // This file does not exist
	if err != nil {
		fmt.Println("Error expected, but calling f.Name() will cause a panic!")
	}
	// Before Go 1.25: f.Name() could fail silently
	// Go 1.25+: If f is nil, f.Name() will cause a panic
	fmt.Println(f.Name()) // A panic will occur here
}

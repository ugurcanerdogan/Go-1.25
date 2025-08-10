package main

import (
	"crypto/sha3"
	"fmt"
	"reflect"
)

func main() {
	// create hash
	h1 := sha3.New256()
	h2 := sha3.New256()
	// absorbs data
	h1.Write([]byte("hello"))
	h2.Write([]byte("hello"))

	clone, _ := h1.Clone()
	h3 := clone.(*sha3.SHA3)

	// h3 has the same state as h1 and h2, so it will produce the same hash after writing the same data.
	h1.Write([]byte("world"))
	h2.Write([]byte("world"))
	h3.Write([]byte("world"))

	fmt.Printf("h1: %x\n", h1.Sum(nil))
	fmt.Printf("h2: %x\n", h2.Sum(nil))
	fmt.Printf("h3: %x\n", h2.Sum(nil))
	fmt.Printf("h1 == h2: %t\n", reflect.DeepEqual(h1, h2))
	fmt.Printf("h2 == h3: %t\n", reflect.DeepEqual(h2, h3))
	fmt.Printf("h1 == h3: %t\n", reflect.DeepEqual(h1, h3))
}

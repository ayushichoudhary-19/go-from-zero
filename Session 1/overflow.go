package main

import "fmt"

func main() {
	var maxUint32 uint32 = 4294967295
	fmt.Println("Max uint32:", maxUint32)
	fmt.Println("After overflow:", maxUint32+1)

	var big uint64 = 4294967296
	fmt.Println("Using uint64 to avoid overflow:", big)
}

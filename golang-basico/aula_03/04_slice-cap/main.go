package main

import "fmt"

func main() {
	// inicializando um slice vazio
	var slice []int

	// crescendo dinamicamente com append
	for i := 1; i < 10; i++ {
		slice = append(slice, i)
		fmt.Printf("Após append %d: Slice: %v, Len: %d, Cap: %d\n", i, slice, len(slice), cap(slice))
	}
}

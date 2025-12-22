package main

import "fmt"

func main() {
	// usando make para criar um slice com capacidade inicial
	slice := make([]int, 0, 5) // len 0, cap 5
	fmt.Printf("inicial: slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	// crescendo detro da capacidade inicial
	for i := 1; i <= 5; i++ {
		slice = append(slice, i)
		fmt.Printf("após append: slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	}

	// crescendo além da capacidade inicial
	slice = append(slice, 6)
	fmt.Printf("após append 6: slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
}

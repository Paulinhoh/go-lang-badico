package main

import "fmt"

func main() {
	// criando um array
	array := [5]int{10, 20, 30, 40, 50}

	// criando um slice basedo no array
	slice := array[1:4] //inclui indices 1,2 e 3
	fmt.Println("array:", array)
	fmt.Printf("Slice: %v, len: %d, cap:%d\n", slice, len(slice), cap(slice))

	// modificar o slice afeta o array subjacente
	slice[0] = 100
	fmt.Println("após modificar slice:")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}

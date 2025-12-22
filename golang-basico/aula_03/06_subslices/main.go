package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}
	subslice := original[1:4] // sub-slice: elemento de indice 1 a 3
	fmt.Println("original:", original)
	fmt.Println("subslice:", subslice)

	// modificando o sub-slice
	subslice[0] = 99
	subslice = append(subslice, 44)

	fmt.Println("Após modificar sub-slice:")
	fmt.Println("original:", original)
	fmt.Println("subslice:", subslice)
}

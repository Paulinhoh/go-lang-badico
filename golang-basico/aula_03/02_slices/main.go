package main

import "fmt"

func main() {
	slice := []int{10, 20, 30} // lista dinamica (array dinâmico)
	slice = append(slice, 40, 50)

	fmt.Println(slice)
}

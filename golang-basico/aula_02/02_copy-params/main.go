package main

import "fmt"

func alteraCopia(x int) {
	fmt.Println("recebido como:", x)
	x = x * 2
	fmt.Println("atualizado para:", x)
}

func main() {
	numero := 10
	fmt.Println("definido como:", numero)
	alteraCopia(numero)
	fmt.Println("fora da função:", numero)
}

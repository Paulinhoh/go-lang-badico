package main

import "fmt"

func alteraOriginal(x *int) {
	fmt.Println("recebido como:", *x)
	*x = *x * 2
	fmt.Println("atualizado para:", *x)
}

func main() {
	numero := 10
	fmt.Println("definido como:", numero) 
	alteraOriginal(&numero)
	fmt.Println("fora da função:", numero)
}

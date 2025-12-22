package main

import "fmt"

func main() {
	fmt.Println("Hello from go world! Lets build")

	var name string
	fmt.Print("Qual o seu nome: ")
	fmt.Scanln(&name)

	fmt.Printf("Bem-vindo ao mundo de Go, %s!\n", name)
}

// go build -o myapp main.go -> comando para buildar o arquivo main.go e gerar um executável chamado myapp
// ./myapp -> comando para executar o arquivo gerado myapp
// go run main.go -> comando para compilar e executar o arquivo main.go em um único passo

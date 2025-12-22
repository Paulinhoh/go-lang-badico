package main

import (
	"fmt"
	"os"
)

// exemplo de como é implementado o io-write pelo go
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ------------------------------------------------

func main() {
	// criar ou abrir um arquivo para escrita
	file, err := os.Create("01_io-writer/example.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close() // certifique-se de fechar o arquivo no final

	// dados a serem gravados
	data := []byte("Olá, mundo!")

	// escrevendo dados no arquivo
	n, err := file.Write(data)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Printf("Gravados %d bytes no arquivo.\n", n) 
}

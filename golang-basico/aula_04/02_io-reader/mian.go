package main

import (
	"fmt"
	"os"
)

// exemplo de como é implementado o io-write pelo go
type Reader interface {
	Read(p []byte) (n int, err error)
}

// ------------------------------------------------

func main() {
	// abrir um arquivo para leitura
	file, err := os.Open("02_io-reader/example.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// criar um buffer para armazenar os dados lidos
	buffer := make([]byte, 1024)

	// lendo os dados do arquivo
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// convertendo os dados lidos para string e exibindo
	fmt.Printf("Lido %d bytes: %s\n", n, string(buffer[:n]))
}

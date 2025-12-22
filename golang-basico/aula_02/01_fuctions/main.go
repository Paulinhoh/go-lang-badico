package main

import "fmt"

func dobrar(numero int) (int, error) {
	if numero < 0 {
		return 0, fmt.Errorf("numero negativo: %d", numero)
	}

	return numero * 2, nil
}

// string -> ""
// int/float -> 0
// bool -> false
// endereço de memoria -> nil

func main() {
	resultado, err := dobrar(-10)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", resultado)
}

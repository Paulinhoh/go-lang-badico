package main

import "fmt"

func Main1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Esta eh a iteração numero %d\n", i)
	}

	// outra forma de fazer o msm for
	for i := range 5 {
		fmt.Printf("Esta eh a iteração numero %d\n", i)
	}
}

func main() {
	var age int

	for {
		fmt.Printf("Por favor diga a sua idade: (Obrigatorio ser maior que 18): ")
		_, err := fmt.Scanln(&age)
		if err != nil {
			fmt.Println("Entrada invalida, por favor repita")
			continue
		}

		if age < 18 {
			fmt.Println("Entrada invalida, por favor repita")
			continue
		}

		break
	}

	fmt.Println("Muito bem, seja bem-vindo")
}

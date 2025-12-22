package main

import (
	"fmt"
	"time"
)

func get_idade() int {
	var age int

	for {
		fmt.Printf("Qual a sua idade: ")
		_, err := fmt.Scanln(&age)
		if err != nil {
			println("Entrada invalida! Tente novamente")
			continue
		}
		break
	}

	return age
}

func main() {
	age := get_idade()
	date := time.Now().Year()

	dateOfBirth := date - age

	fmt.Println("Voce nasceu em", dateOfBirth)
}

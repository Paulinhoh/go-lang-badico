package main

import "fmt"

func main() {
	var age int
	fmt.Println("Qual a sua idade")
	fmt.Scanln(&age)

	if age > 18 {
		println("Vc é de maior meu chapa")
	} else {
		fmt.Println("vc é de menor meu amigo")
	}

	if len("") > 0 {
		fmt.Println("go n tem truth")
	}
}

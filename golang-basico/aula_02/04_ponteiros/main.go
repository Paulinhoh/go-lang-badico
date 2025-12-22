package main

import "fmt"

func main() {
	// uma variavel normal
	idade := 30

	// um ponteiro para a variavel `idade`
	ponteiroIdade := &idade

	fmt.Println("valor de idade:", idade)              // -> 30
	fmt.Println("endereço de idade:", ponteiroIdade)   // -> mostra o endereço de memoria
	fmt.Println("valor via ponteiro:", *ponteiroIdade) // -> 30 (desreferencia)

	// ------------------------------------------------------------------------------

	numero := 42
	ponteiro := &numero

	fmt.Println("\nantes da alteração:", numero)
	*ponteiro = 99
	fmt.Println("depois da alteração:", numero)

}

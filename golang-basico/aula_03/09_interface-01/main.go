package main

import "fmt"

type Carro struct {
	Modelo string
}

func (c Carro) Acelerar() {
	fmt.Printf("O carro %s esta acelerando!\n", c.Modelo)
}

type Bicicleta struct {
	Tipo string
}

func (b Bicicleta) Acelerar() {
	fmt.Printf("A bicicleta  do tipo %s esta acelerando!\n", b.Tipo)
}

// definindo uma interface para comportamentos
type Veiculo interface {
	Acelerar()
}

func TestarVeiculo(v Veiculo) {
	v.Acelerar()
}

func main() {
	carro := Carro{Modelo: "Sedan"}
	bike := Bicicleta{Tipo: "Mountain bike"}
	TestarVeiculo(carro)
	TestarVeiculo(bike)
}
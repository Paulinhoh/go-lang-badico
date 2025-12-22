package main

import "fmt"

// interfaces são sempre ponteiros
// comportamentos com interfaces
type Acelerador interface {
	Acelerar()
}

type Freio interface {
	Freiar()
}

// struct base: combinando comportamentos
type Carro struct {
	Modelo string
}

func (c Carro) Acelerar() {
	fmt.Printf("O carro %s esta acelerando!\n", c.Modelo)
}

func (c Carro) Freiar() {
	fmt.Printf("O carro %s esta freiando!\n", c.Modelo)
}

// adicionando novos comportamentos por composição
type CarroEletrico struct {
	Carro
	BateriaCarga int
}

func (ce CarroEletrico) carregarBateria() {
	fmt.Printf("Carregando bateria... Carga: %d%%\n", ce.BateriaCarga)
}

func main() {
	ce := CarroEletrico{
		Carro:        Carro{Modelo: "Tesla"},
		BateriaCarga: 80,
	}

	ce.Acelerar()        // metodo herdado por composição
	ce.Freiar()          // metodo herdado por composição
	ce.carregarBateria() // metodo do carro eletrico
}

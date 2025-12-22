package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

// metodo para apresentar uma Pessoa
func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Olá, eu sou %s e tenho %d anos.", p.Nome, p.Idade)
}

// metodo que modifica a struct Pessoa
func (p *Pessoa) Envelhecer() {
	p.Idade += 1
}

func main() {
	pessoa := Pessoa{"Paulo Henrique", 24}

	fmt.Println("Nome:", pessoa.Nome)
	fmt.Println("Idade:", pessoa.Idade)
	fmt.Println("Pessoa:", pessoa)
	fmt.Println("Pessoa &:", &pessoa)
	fmt.Println("Pessoa.Nome &:", &pessoa.Nome)
	fmt.Println("Pessoa.Idade &:", &pessoa.Idade)

	fmt.Println("\nPaulinho diz:", pessoa.Apresentar())
	pessoa.Envelhecer()
	fmt.Println("Paulinho agora diz:", pessoa.Apresentar())
}

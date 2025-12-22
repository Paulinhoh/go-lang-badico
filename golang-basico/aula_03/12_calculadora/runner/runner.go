package runner

import (
	"fmt"

	"github.com/Paulinhoh/calculadora/operacao"
)

type Operacao interface {
	Calcular(a, b float64) float64
}

type Runner struct {
	Operacoes map[string]Operacao
}

func (r *Runner) Executar() {
	var a, b float64
	var operacao string
	fmt.Print("Digite o primeiro numero: ")
	fmt.Scanln(&a)
	fmt.Print("Digite o segundo numero: ")
	fmt.Scanln(&b)
	fmt.Print("Escolha a operação (+ - * /): ")
	fmt.Scanln(&operacao)

	op, existe := r.Operacoes[operacao]
	if !existe {
		fmt.Println("Operação invalida")
		return
	}

	resultado := op.Calcular(a, b)
	fmt.Printf("Resultado: %.2f\n", resultado)
}

func NewRunner() *Runner {
	return &Runner{
		Operacoes: map[string]Operacao{
			"+": operacao.Soma{},
			"-": operacao.Subtracao{},
			"*": operacao.Multiplicacao{},
			"/": operacao.Divisao{},
		},
	}
}

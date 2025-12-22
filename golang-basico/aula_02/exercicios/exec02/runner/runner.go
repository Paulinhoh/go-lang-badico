package runner

import (
	"errors"
	"fmt"
	"golang-basico/calculadora"
)

type Runner struct {
	calculadora *calculadora.Calculadora
	operacao    string
}

// metodo para solicitar valores
func (r *Runner) SolicitarValores() error {
	fmt.Print("Digite o primeiro númeor para iniciar: ")
	if _, err := fmt.Scanln(&r.calculadora.Operando1); err != nil {
		return errors.New("entrada invalida para o primeiro numero")
	}

	fmt.Print("Digite o segundo número: ")
	if _, err := fmt.Scanln(&r.calculadora.Operando2); err != nil {
		return errors.New("entrada invalida para o segundo numero")
	}

	return nil
}

// metodo para solicitar operação
func (r *Runner) SolicitarOperacao() error {
	var operacao string
	fmt.Print("Escolha a operação (+ - * /): ")
	if _, err := fmt.Scanln(&operacao); err != nil {
		return errors.New("entrada invalida para a operação.")
	}

	// validação de operação
	switch operacao {
	case "+", "-", "*", "/":
		r.operacao = operacao
		return nil
	default:
		return errors.New("operação invalida!")
	}
}

// metodo para executar a operação
func (r *Runner) ExecutarOperacao() {
	switch r.operacao {
	case "+":
		fmt.Println("Resultado:", r.calculadora.Somar())
	case "-":
		fmt.Println("Resultado:", r.calculadora.Subtrair())
	case "*":
		fmt.Println("Resultado:", r.calculadora.Multiplicar())
	case "/":
		resultado, err := r.calculadora.Dividir()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Resultado:", resultado)
	default:
		fmt.Println("Operação inválida!")
	}
}

func (r *Runner) Execute() {
	for {
		// solicitar valores com validação
		if err := r.SolicitarValores(); err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// solicitar operação com validação
		if err := r.SolicitarOperacao(); err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// executar a operação escolhida
		r.ExecutarOperacao()
		break
	}
}

func NewRunner(c *calculadora.Calculadora) *Runner {
	return &Runner{calculadora: c}
}

package main

import "fmt"

type Calculadora struct {
	Operando1 float64
	Operando2 float64
}

func (c Calculadora) Dividir() (float64, error) {
	if c.Operando2 == 0 {
		return 0, fmt.Errorf("error: divisão por zero.")
	}
	return c.Operando1 / c.Operando2, nil
}

func (c Calculadora) Somar() float64 {
	return c.Operando1 + c.Operando2
}

func (c Calculadora) Subtrair() float64 {
	return c.Operando1 - c.Operando2
}

func (c Calculadora) Multiplicar() float64 {
	return c.Operando1 * c.Operando2
}

func main() {
	var operacao string
	var op1, op2 float64

	fmt.Printf("Digite o valor 1: ")
	fmt.Scanln(&op1)
	fmt.Printf("Digite o valor 2: ")
	fmt.Scanln(&op2)
	fmt.Printf("Digite a operação (+ - * /): ")
	fmt.Scanln(&operacao)
	calc := Calculadora{op1, op2}

	switch operacao {
	case "+":
		fmt.Println("Resultado:", calc.Somar())
	case "-":
		fmt.Println("Resultado:", calc.Subtrair())
	case "*":
		fmt.Println("Resultado:", calc.Multiplicar())
	case "/":
		resultado, err := calc.Dividir()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resultado)
	default:
		fmt.Println("Operação inválida!")
	}

}

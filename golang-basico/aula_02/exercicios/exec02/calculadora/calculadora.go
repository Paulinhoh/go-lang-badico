package calculadora

import "errors"

type Calculadora struct {
	Operando1 float64
	Operando2 float64
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

func (c Calculadora) Dividir() (float64, error) {
	if c.Operando2 == 0 {
		return 0, errors.New("divisão por zero.")
	}
	return c.Operando1 / c.Operando2, nil
}

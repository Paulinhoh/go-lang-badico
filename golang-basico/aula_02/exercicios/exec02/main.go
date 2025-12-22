package main

import (
	"golang-basico/calculadora"
	"golang-basico/runner"
)

func main() {
	calc := &calculadora.Calculadora{}
	runner := runner.NewRunner(calc)
	runner.Execute()
}

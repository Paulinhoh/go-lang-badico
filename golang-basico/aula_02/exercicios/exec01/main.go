package exec01

func main() {
	calc := &Calculadora{}
	runner := NewRunner(calc)
	runner.Execute()
}

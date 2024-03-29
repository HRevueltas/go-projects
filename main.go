package main

import "fmt"

type Transaction struct {
	Amount float64
}

func (t *Transaction) CalculateTax(rate float64) float64 {
	return t.Amount * rate
}

func main() {
	tr := &Transaction{Amount: 100.0}
	tax := tr.CalculateTax(0.15)
	fmt.Println("Tax:", tax)
}

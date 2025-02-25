package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	desc := `=================
Profit Calculator
=================`

	fmt.Println(desc)

	fmt.Print("Revenue: ")
	_, _ = fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	_, _ = fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	_, _ = fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateResult(revenue, expenses, taxRate)
	outputResult(ebt, profit, ratio)
}

func calculateResult(revenue float64, expenses float64, taxRate float64) (earningBeforeTax float64, earningAfterTax float64, ratio float64) {
	earningBeforeTax = revenue - expenses
	earningAfterTax = earningBeforeTax * (1 - taxRate/100)
	ratio = earningBeforeTax / earningAfterTax
	return
}

func outputResult(earningBeforeTax float64, earningAfterTax float64, ratio float64) {
	fmt.Printf("Earning before tax: %.2f\n", earningBeforeTax)
	fmt.Printf("Earning after tax (profit): %.2f\n", earningAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

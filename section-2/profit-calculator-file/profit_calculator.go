package main

import (
	"fmt"
	"os"
)

const profitFile = "profit.txt"

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

	if revenue < 0 {
		fmt.Println("Revenue cannot be 0")
		return
	}

	fmt.Print("Expenses: ")
	_, _ = fmt.Scan(&expenses)

	if expenses < 0 {
		fmt.Println("Expenses cannot be 0")
		return
	}

	fmt.Print("Tax Rate: ")
	_, _ = fmt.Scan(&taxRate)

	if taxRate < 0 || taxRate > 100 {
		fmt.Println("Tax rate must be between 0 and 100")
		return
	}

	ebt, profit, ratio := calculateResult(revenue, expenses, taxRate)
	outputResult(ebt, profit, ratio)

	err := writeResultToFile(fmt.Sprintf("%.2f;%.2f;%.2f", ebt, profit, ratio), profitFile)
	if err != nil {
		fmt.Println(err)
	}
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

func writeResultToFile(result string, filename string) error {
	return os.WriteFile(filename, []byte(result), 0644)
}

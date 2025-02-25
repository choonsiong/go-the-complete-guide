package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter investment amount: ")
	_, _ = fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	_, _ = fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	_, _ = fmt.Scan(&years)

	fmt.Println("Inflation rate: ", inflationRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Real Value: %.2f\n", futureRealValue)
}

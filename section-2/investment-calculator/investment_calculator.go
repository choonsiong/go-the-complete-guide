package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5

	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Real Value:", futureRealValue)
}

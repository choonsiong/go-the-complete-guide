package main

import (
	"example.com/price-calc-v3/filemanager"
	"example.com/price-calc-v3/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := job.Process("prices.txt", "result")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("tax rate: $%.2f\n", taxRate)
		fmt.Println(job.TaxIncludedPrices.PrettyFormat())
	}

	// This is for demo only, it doesn't make sense to have to enter
	// the same inputs from stdin for each tax rate...
	//for _, taxRate := range taxRates {
	//	cm := cmdmanager.New()
	//	job := prices.NewTaxIncludedPriceJob(cm, taxRate)
	//	err := job.Process("prices.txt", "result")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	fmt.Printf("tax rate: $%.2f\n", taxRate)
	//	fmt.Println(job.TaxIncludedPrices.PrettyFormat())
	//}
}

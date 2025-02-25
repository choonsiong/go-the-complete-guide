package main

import (
	"example.com/price-calc-channel/filemanager"
	"example.com/price-calc-channel/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	dones := make([]chan *prices.TaxIncludedPriceJob, len(taxRates))
	//dones := make([]chan bool, len(taxRates))
	errors := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		//dones[i] = make(chan bool)
		dones[i] = make(chan *prices.TaxIncludedPriceJob)
		errors[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go job.Process("prices.txt", dones[i], errors[i])
		//fmt.Printf("tax rate: $%.2f\n", taxRate)
		//fmt.Println(job.TaxIncludedPrices.PrettyFormat())
	}

	// Below will not work, since we are waiting for something
	// to send on the error channel
	//for _, e := range errors {
	//	<-e
	//}
	//
	//for _, ch := range dones {
	//	<-ch
	//}

	//for i, _:= range taxRates {
	for i, rate := range taxRates {
		select {
		case err := <-errors[i]:
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(err)
		//case <-dones[i]:
		//	fmt.Println("done")
		case job := <-dones[i]:
			fmt.Printf("tax rate: $%.2f\n", rate)
			fmt.Println(job.TaxIncludedPrices.PrettyFormat())
		}
	}
}

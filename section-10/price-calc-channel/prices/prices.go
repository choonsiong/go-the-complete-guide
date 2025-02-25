package prices

import (
	"example.com/price-calc-channel/conversion"
	"example.com/price-calc-channel/iomanager"
	"fmt"
	"strconv"
)

type TaxIncludedPricesMap map[string]float64

func (tm *TaxIncludedPricesMap) PrettyFormat() string {
	result := ""
	for k, v := range *tm {
		result += fmt.Sprintf("price: $%s, price included tax: $%.2f\n", k, v)
	}
	return result
}

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager  `json:"-"`
	TaxRate           float64              `json:"tax_rate"`
	InputPrices       []float64            `json:"input_prices"`
	TaxIncludedPrices TaxIncludedPricesMap `json:"tax_included_prices"`
}

// NewTaxIncludedPriceJob returns a new *TaxIncludedPriceJob
func NewTaxIncludedPriceJob(m iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: m,
		TaxRate:   taxRate,
	}
}

// Process processes the file given by name and writes the result to a JSON
// file prefixes with the given outPrefix
// func (t *TaxIncludedPriceJob) Process(name string, ch chan bool, errch chan error) {
func (t *TaxIncludedPriceJob) Process(name string, ch chan *TaxIncludedPriceJob, errch chan error) {
	err := t.loadPricesFromFile(name)
	if err != nil {
		errch <- err
		return
	}

	result := make(map[string]float64)

	for _, price := range t.InputPrices {
		taxIncludedPrice := price * (1 + t.TaxRate)
		taxIncludedPriceStr := fmt.Sprintf("%.2f", taxIncludedPrice)
		taxIncludedPriceFloat, err := strconv.ParseFloat(taxIncludedPriceStr, 64)
		if err != nil {
			errch <- err
			return
		}
		result[fmt.Sprintf("%.2f", price)] = taxIncludedPriceFloat
	}

	t.TaxIncludedPrices = result

	err = t.IOManager.WriteResult(t)
	if err != nil {
		errch <- err
		return
	}

	// Instead of returning bool, now we return the struct itself
	//ch <- true
	ch <- t
}

// loadPricesFromFile loads all the prices from file into the TaxIncludedPriceJob.InputPrices
// field and returns an error if any
func (t *TaxIncludedPriceJob) loadPricesFromFile(name string) error {
	pricesInString, err := t.IOManager.ReadLines()
	if err != nil {
		return err
	}

	pricesInFloat, err := conversion.StringsToFloats(pricesInString)
	if err != nil {
		return err
	}

	t.InputPrices = pricesInFloat

	return nil
}

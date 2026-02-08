package prices

import (
	"fmt"

	"example.com/taxes/conversions"
	"example.com/taxes/iomanager"
)

type TaxPriceJob struct {
	TaxRate   float64             `json:"tax_rate"`
	Prices    []float64           `json:"prices"`
	TaxPrices map[string]string   `json:"tax_prices"`
	IOManager iomanager.IOManager `json:"-"`
}

func (job *TaxPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxInclPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxInclPrice)
	}

	fmt.Println(result)

	job.TaxPrices = result

	return job.IOManager.WriteResult(job)
}

func (job *TaxPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversions.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.Prices = prices
	return nil
}

func NewTaxPriceJob(iom iomanager.IOManager, taxRate float64) *TaxPriceJob {
	return &TaxPriceJob{
		TaxRate:   taxRate,
		IOManager: iom,
	}
}

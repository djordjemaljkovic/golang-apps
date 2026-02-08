package main

import (
	"fmt"

	"example.com/taxes/cmdmanager"
	"example.com/taxes/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, rate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		cmd := cmdmanager.New()
		priceJob := prices.NewTaxPriceJob(cmd, rate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

	fmt.Println(result)
}

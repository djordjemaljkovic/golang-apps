package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, revenueErr := getInput("Revenue")

	expenses, expErr := getInput("Expenses")

	taxRate, taxErr := getInput("Tax Rate")

	if taxErr != nil || expErr != nil || revenueErr != nil {
		fmt.Println(taxErr, expErr, revenueErr)
		return
		// panic(err)
	}

	ebt, profit, ratio := calcFinancials(revenue, expenses, taxRate)

	fmt.Printf(" Ebit: %v\n Profit: %v\n Ratio: %.2f", ebt, profit, ratio)

	storeResults(ebt, ratio, profit)
}

func getInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text, ": ")
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number")
	}

	return userInput, nil
}

func calcFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, ratio, profit float64) {
	results := fmt.Sprintf(" Ebit: %.1f\n Profit: %.1f\n Ratio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

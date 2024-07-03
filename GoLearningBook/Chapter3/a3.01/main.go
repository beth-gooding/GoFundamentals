package main

import "fmt"

func salesTaxCalc(cost float64, rate float64) float64 {
	return cost * rate
}

func main() {
	var totalSalesTax float64 = 0
	totalSalesTax += salesTaxCalc(0.99, 0.075)
	totalSalesTax += salesTaxCalc(2.75, 0.015)
	totalSalesTax += salesTaxCalc(0.87, 0.02)

	fmt.Println("Sales Tax Total:", totalSalesTax)
}

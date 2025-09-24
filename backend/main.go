package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to Best Mortgage!")

	var loan_amount float64 = 158000
	var term_years float64 = 30
	var interest_rate float64 = 4.5
	result := calcMortgage(loan_amount, term_years, interest_rate)

	fmt.Printf("Monthy Repayments are %v \n", result)
}

func calcMortgage(loan_amount float64, term_years float64, interest_rate float64) float64 {
	n := term_years * 12 // number of monthly payments
	r := interest_rate / 12 / 100
	result := loan_amount * r * math.Pow(1+r, n) / (math.Pow(1+r, n) - 1)
	result = math.Round(result*100) / 100
	return result
}

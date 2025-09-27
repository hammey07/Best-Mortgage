package main

import (
	"best-mortgage/providers/banks"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// centralbank.PrintCentralBankRates()
	bankRates, err := banks.LoadRates()
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, bank := range bankRates.Banks {
		fmt.Println(bank.BankName, bank.Fixed1To3Green)
	}

	// handleRequests()

	// var loan_amount float64 = 158000
	// var term_years float64 = 30
	// var interest_rate float64 = 4.5
	// result := calcMortgage(loan_amount, term_years, interest_rate)

	// fmt.Printf("Monthy Repayments are %v \n", result)

	// func calcMortgage(loan_amount float64, term_years float64, interest_rate float64) float64 {
	// 	n := term_years * 12 // number of monthly payments
	// 	r := interest_rate / 12 / 100
	// 	result := loan_amount * r * math.Pow(1+r, n) / (math.Pow(1+r, n) - 1)
	// 	result = math.Round(result*100) / 100
	// 	return result
	// }
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Best Mortgage!")
	fmt.Fprint(w, "Available Endpoints")
}

func getRates(w http.ResponseWriter, r *http.Request) {
	type Rate struct {
		Type         string  `json:title`
		Provider     string  `json:provider`
		Duration     float64 `json:duration`
		InterestRate float64 `json:interestRate`
	}

	type Rates []Rate

	rates := Rates{
		Rate{Type: "Fixed",
			Provider:     "AIB",
			Duration:     3,
			InterestRate: 4.5},
		Rate{Type: "Fixed",
			Provider:     "PTSB",
			Duration:     5,
			InterestRate: 4.35},
		Rate{Type: "Variable",
			Provider:     "Revolut",
			Duration:     2.5,
			InterestRate: 4.15},
	}

	fmt.Println("Endpoint: Rates endpoint")
	json.NewEncoder(w).Encode(rates)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/rates", getRates)
	// http.HandleFunc("/centralbank/latest", )
	log.Fatal(http.ListenAndServe(":8001", nil))
}

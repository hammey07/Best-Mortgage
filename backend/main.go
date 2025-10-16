package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// centralbank.PrintCentralBankRates()
	handleRequests()

}

type Rates struct {
	New         map[string]float64 `json:"new"`
	Outstanding map[string]float64 `json:"outstanding"`
	Green       map[string]float64 `json:"green"`
	Variable    map[string]float64 `json:"variable"`
}

type Bank struct {
	BankName string `json:"bank_name"`
	Rates    Rates  `json:"rates"`
}

type BankRates struct {
	ReportDate string `json:"reporting_date"`
	Banks      []Bank `json:"banks"`
}

type BankResponse struct {
	BankName string                        `json:"bank_name`
	Rates    map[string]map[string]float64 `json:"rates`
}

type Response struct {
	ReportDate string         `json:"reporting_date`
	Banks      []BankResponse `json:"banks"`
}

func LoadRates() (BankRates, error) {
	data, err := os.ReadFile("banks/data.json")
	if err != nil {
		fmt.Println("Error reading data", err)
		return BankRates{}, err
	}

	var rates BankRates
	err = json.Unmarshal(data, &rates)
	if err != nil {
		fmt.Println("Error parsing json", err)
		return BankRates{}, err

	}
	return rates, nil
}

func BuildResponse(bankRates BankRates, userType string) Response {
	resp := Response{
		ReportDate: bankRates.ReportDate,
		Banks:      []BankResponse{},
	}

	for _, bank := range bankRates.Banks {
		b := BankResponse{
			BankName: bank.BankName,
			Rates:    map[string]map[string]float64{},
		}

		switch userType {
		case "first_time_buyer":
			b.Rates["new"] = bank.Rates.New
			b.Rates["green"] = bank.Rates.Green
			b.Rates["variable"] = bank.Rates.Variable
		case "remortgage":
			b.Rates["outstanding"] = bank.Rates.Outstanding
			b.Rates["variable"] = bank.Rates.Variable
		}

		resp.Banks = append(resp.Banks, b)
	}
	return resp
}

func handleRequests() {
	bankRates, err := LoadRates()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	http.HandleFunc("/", homePage)
	http.HandleFunc("/rates/first_time_buyer", func(w http.ResponseWriter, r *http.Request) {
		response := BuildResponse(bankRates, "first_time_buyer")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	http.HandleFunc("/rates/remortgage", func(w http.ResponseWriter, r *http.Request) {
		response := BuildResponse(bankRates, "remortgage")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	// http.HandleFunc("/centralbank/latest", )
	log.Fatal(http.ListenAndServe(":8001", nil))
}

// var loan_amount float64 = 158000
// var term_years float 64 = 30
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

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Welcome to Smart Mortgage. This project is currently in development. The following endpoints provides test data only.<br>")
	fmt.Fprint(w, `<br>Available Endpoints<br>`)
	fmt.Fprint(w, `First Time Buyer Rates : <a href="http://localhost:8001/rates/first_time_buyer">First Time Buyer</a><br>`)
	fmt.Fprint(w, `Remortgage Rates : <a href="http://localhost:8001/rates/remortgage">Remortgage </a><br>`)

}

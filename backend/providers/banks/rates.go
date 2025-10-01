package banks

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func LoadRates() (BankRates, error) {
	// dir, _ := os.Getwd()
	// fmt.Println("Current working directory:", dir)
	data, err := os.ReadFile("providers/banks/data.json")
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
	// for _, bank := range rates.Banks {
	// 	fmt.Println(bank)
	// 	fmt.Println(bank.Fixed1To3Green)
	// }
	// fmt.Println("Data loaded:", rates.ReportDate)
	return rates, nil
}

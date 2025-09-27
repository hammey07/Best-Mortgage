package banks

import (
	"encoding/json"
	"fmt"
	"os"
)

type Bank struct {
	BankName              string  `json:"bank_name"`
	Fixed1To3Outstanding  float64 `json:"fixed_1_3_outstanding"`
	Fixed1To3New          float64 `json:"fixed_1_3_new"`
	Fixed1To3Green        float64 `json:"fixed_1_3_green"`
	FixedOver3Outstanding float64 `json:"fixed_over_3_outstanding"`
	FixedOver3New         float64 `json:"fixed_over_3_new"`
	FixedOver3Green       float64 `json:"fixed_over_3_green"`
	VariableStandard      float64 `json:"variable_standard"`
	VariableTracker       float64 `json:"variable_tracker"`
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

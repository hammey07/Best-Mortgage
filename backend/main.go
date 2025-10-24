package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

type Mortgage struct {
	Lender         string `json:"lender"`
	Product        string `json:"product"`
	Rate           string `json:"rate"`
	APRC           string `json:"APRC"`
	MaxLTV         string `json:"max_LTV"`
	MonthlyPayment string `json:"monthly_payment"`
}

type MortgageData struct {
	FirstTimeBuyer []Mortgage `json:"first_time_buyer_mortgages"`
	Remortgage     []Mortgage `json:"remortgage_mortgages"`
}

func loadData(filePath string) (MortgageData, error) {
	var data MortgageData
	file, err := os.Open(filePath)
	if err != nil {
		return data, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	return data, err
}

func main() {
	data, err := loadData("banks/data.json") // Your JSON file path
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	http.HandleFunc("/rates/first-time-buyer", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data.FirstTimeBuyer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/rates/switcher", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data.Remortgage); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "Welcome to Smart Mortgage. This project is currently in development. The following endpoints provides test data only.<br>")
		fmt.Fprint(w, `<br>Available Endpoints<br>`)
		fmt.Fprint(w, `First Time Buyer Rates : <a href="http://localhost:8001/rates/first-time-buyer">First Time Buyer</a><br>`)
		fmt.Fprint(w, `Remortgage Rates : <a href="http://localhost:8001/rates/switcher">Remortgage </a><br>`)

	})

	log.Println("Server running at http://localhost:8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}

package centralbank

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Record struct {
	ReportingDate string  `json:"Reporting date"`
	FO_1_3        float64 `json:"principal_dwelling_houses__fixed_rate__from_1_to_3_years__ra"`
	FO_3_X        float64 `json:"principal_dwelling_houses__fixed_rate__over_3_years__rates_o"` // outstanding
	FNB_1_3       float64 `json:"principal_dwelling_houses__fixed_rate__from_1_to_3_years__ra_2"`
	FNB_3X        float64 `json:"principal_dwelling_houses__fixed_rate__over_3_years__rates_o_2"` // outstanding
}

type APIResponse struct {
	Result struct {
		Records []Record `json:"records"`
	} `json:"result"`
}

func FetchLatestRates() (*Record, error) {
	// * is a pointer towards struct
	query := "https://opendata.centralbank.ie/api/3/action/datastore_search?resource_id=fb07a41b-15f7-4697-a7f7-ce8209d794e5&limit=10000"
	resp, err := http.Get(query) // sending get request
	if err != nil {
		fmt.Print("Error making request", err)
		return nil, err
	}
	defer resp.Body.Close() //ensureing connection is closed after we are done

	body, err := io.ReadAll(resp.Body) // sending get request
	if err != nil {
		fmt.Print("Error with reading response", err)
		return nil, err
	}

	var data APIResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Print("Error with parsing json", err)
		return nil, err
	}

	if len(data.Result.Records) == 0 {
		fmt.Print("No records json", err)
		return nil, err
	}
	records := data.Result.Records
	latest := records[len(records)-1]
	return &latest, nil
}

func PrintCentralBankRates() {
	latest, err := FetchLatestRates()
	if err != nil {
		println("Error fetching rates", err)
		return
	}

	fmt.Println("Latest Central Bank Rates")
	fmt.Println("Principal Dwelling Houses, fixed rate, from 1 to 3 years, rates on outstanding amounts (%)")
	fmt.Println(latest.FO_1_3)
	fmt.Println("Principal Dwelling Houses, fixed rate, over 3 years, rates on outstanding amounts (%)")
	fmt.Println(latest.FO_3_X)
	fmt.Println("Principal Dwelling Houses, fixed rate, from 1 to 3 years, rates on new business (%)")
	fmt.Println(latest.FNB_1_3)
	fmt.Println(" Principal Dwelling Houses, fixed rate, over 3 years, rates on new business (%)")
	fmt.Println(latest.FNB_3X)

}

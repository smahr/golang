package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type CalculationResult struct {
	Result float64
}

type CalculationError struct {
	Result string
}

func caculateDonationCredit(usd float64) float64 {
	if usd == 0 {
		return 0
	}
	if usd < 10 {
		return 0.1
	} else if usd < 100 {
		return 0.15
	} else if usd < 1000 {
		return 0.2
	} else if usd < 10000 {
		return 0.5
	} else {
		return 2
	}
}

func Calculate(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET params were:", r.URL.Query())
	salesAmountString := r.URL.Query().Get("salesAmount")

	if salesAmount, err := strconv.ParseFloat(salesAmountString, 64); err == nil {
		fmt.Println("bla")
		fmt.Println("input variable is:", salesAmount)
		donCredits := caculateDonationCredit(salesAmount)

		calculationResult := CalculationResult{Result: donCredits}
		json.NewEncoder(w).Encode(calculationResult)
	}

}

func calculationHandler(w http.ResponseWriter, r *http.Request) {
	Calculate(w, r)
}

func setupRoutes() {
	http.HandleFunc("/", calculationHandler)
}

func main() {
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}

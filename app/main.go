package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CalculationResult struct {
	Result float64
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	calculationResult := CalculationResult{Result: 1234.0}

	json.NewEncoder(w).Encode(calculationResult)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	Calculate(w, r)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}

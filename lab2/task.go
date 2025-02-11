package main

import (
	"html/template"
	"lab2/logic"
	"log"
	"net/http"
	"strconv"
)

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("task.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	bCoal := parseFormValue(r, "coal")
	bFuel := parseFormValue(r, "fuel")
	bGas := parseFormValue(r, "gas")

	qCoal, gVynCoal, aVynCoal, aRCoal, nZuCoal, kHSCoal := 20.47, 1.5, 0.8, 25.2, 0.985, 0.0
	qFuel, gVynFuel, aVynFuel, aRFuel, nZuFuel, kHSFuel := 40.40, 0.0, 1.0, 0.15, 0.985, 0.0
	qGas := 20.47
	kGasValue := 0.0

	kCoalValue := logic.CalculateK(qCoal, aVynCoal, aRCoal, gVynCoal, nZuCoal, kHSCoal)
	eCoalValue := logic.CalculateE(kCoalValue, bCoal, qCoal)
	kFuelValue := logic.CalculateK(qFuel, aVynFuel, aRFuel, gVynFuel, nZuFuel, kHSFuel)
	eFuelValue := logic.CalculateE(kFuelValue, bFuel, qFuel)
	eGasValue := logic.CalculateE(kGasValue, bGas, qGas)

	data := struct {
		Success bool
		KCoal   float64
		ECoal   float64
		KFuel   float64
		EFuel   float64
		KGas    float64
		EGas    float64
	}{
		Success: true,
		KCoal:   logic.Round(kCoalValue),
		ECoal:   logic.Round(eCoalValue),
		KFuel:   logic.Round(kFuelValue),
		EFuel:   logic.Round(eFuelValue),
		KGas:    logic.Round(kGasValue),
		EGas:    logic.Round(eGasValue),
	}

	tmpl.Execute(w, data)
}

func parseFormValue(r *http.Request, key string) float64 {
	valueStr := r.FormValue(key)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Printf("Error parsing %s: %v", key, err)
		return 0
	}
	return value
}

func main() {
	http.HandleFunc("/", handleCalculate)
	http.ListenAndServe(":8000", nil)
}

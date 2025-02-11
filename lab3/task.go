package main

import (
	"html/template"
	"lab3/logic"
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

	PcValue, err1 := strconv.ParseFloat(r.FormValue("Pc"), 64)
	sigmaValue, err2 := strconv.ParseFloat(r.FormValue("sigma"), 64)
	BValue, err3 := strconv.ParseFloat(r.FormValue("B"), 64)

	if err1 != nil || err2 != nil || err3 != nil {
		http.Error(w, "Невірний формат вхідних даних", http.StatusBadRequest)
		return
	}

	profitValue := logic.CalculateP(PcValue, sigmaValue, BValue)

data := struct {
		Success bool
		Profit  float64
	}{
		Success: true,
		Profit:  profitValue,
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

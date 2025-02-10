package main

import (
	"html/template"
	"lab1/logic"
	"log"
	"net/http"
	"strconv"
)

// С - вуглець; Н - водень; S - сірка; N - азот; O - кисень; W - волога; А - зола.
type Data struct {
	C, H, O, S, Q, W, A, V float64
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("task2.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	data := Data{
		C: parseFormValue(r, "C"),
		H: parseFormValue(r, "H"),
		O: parseFormValue(r, "O"),
		S: parseFormValue(r, "S"),
		Q: parseFormValue(r, "Q"),
		W: parseFormValue(r, "W"),
		A: parseFormValue(r, "A"),
		V: parseFormValue(r, "V"),
	}

	dryWorkCoeff := logic.CalculateCoeffDryWork(data.W)
	burnWorkCoeff := logic.CalculateCoeffBurnWork(data.W, data.A)
	workFuelPercentageValue := logic.CalculateWorkFuelPercentage(burnWorkCoeff, dryWorkCoeff, data.C, data.H, data.O, data.S, data.A, data.V)
	heatCombustionValue := data.Q*burnWorkCoeff - 0.025*data.W

	templateData := struct {
		Success                bool
		C, H, O, S, Q, W, A, V float64
		DryWorkCoeff           float64
		BurnWorkCoeff          float64
		WorkFuelPercentage     []float64
		HeatCombustion         float64
	}{
		Success:            true,
		C:                  data.C,
		H:                  data.H,
		O:                  data.O,
		S:                  data.S,
		Q:                  data.Q,
		W:                  data.W,
		A:                  data.A,
		V:                  data.V,
		DryWorkCoeff:       dryWorkCoeff,
		BurnWorkCoeff:      burnWorkCoeff,
		WorkFuelPercentage: workFuelPercentageValue,
		HeatCombustion:     heatCombustionValue,
	}

	tmpl.Execute(w, templateData)
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

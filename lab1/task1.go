package main

import (
	"html/template"
	"lab1/logic"
	"log"
	"net/http"
	"strconv"
)

type CompositionFuelComponents struct {
	H, C, S, N, O, W, A float64
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("task1.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	compositionFuelComponents := CompositionFuelComponents{
		H: parseFormValue(r, "H"),
		C: parseFormValue(r, "C"),
		S: parseFormValue(r, "S"),
		O: parseFormValue(r, "O"),
		N: parseFormValue(r, "N"),
		W: parseFormValue(r, "W"),
		A: parseFormValue(r, "A"),
	}

	coeffWorkDry := logic.CalculateCoeffWorkDry(compositionFuelComponents.W)
	coeffWorkBurn := logic.CalculateCoeffWorkBurn(compositionFuelComponents.W, compositionFuelComponents.A)

	dryCompositionFuelComponents := logic.CalculateNewPercentage(coeffWorkDry, []float64{
		compositionFuelComponents.H,
		compositionFuelComponents.C,
		compositionFuelComponents.S,
		compositionFuelComponents.O,
		compositionFuelComponents.N,
		compositionFuelComponents.A,
	})
	burnCompositionFuelComponents := logic.CalculateNewPercentage(coeffWorkBurn, []float64{
		compositionFuelComponents.H,
		compositionFuelComponents.C,
		compositionFuelComponents.S,
		compositionFuelComponents.O,
		compositionFuelComponents.N,
	})

	heatCombustion := logic.CalculateHeatCombustion(
		compositionFuelComponents.C,
		compositionFuelComponents.H,
		compositionFuelComponents.O,
		compositionFuelComponents.S,
		compositionFuelComponents.W,
	)

	dryFuelHeatCombustion := logic.CalculateDryFuelHeatCombustion(
		heatCombustion,
		compositionFuelComponents.W,
		coeffWorkDry,
	)

	burnFuelHeatCombustion := logic.CalculateBurnFuelHeatCombustion(
		heatCombustion,
		compositionFuelComponents.W,
		coeffWorkBurn,
	)

	// Prepare data for template
	data := struct {
		Success                bool
		InitialComposition     []float64
		DryComposition         []float64
		BurnComposition        []float64
		CoeffWorkDry		   float64
		CoeffWorkBurn		   float64
		HeatCombustion         float64
		DryFuelHeatCombustion  float64
		BurnFuelHeatCombustion float64
	}{
		Success:               true,
		InitialComposition: []float64{
			compositionFuelComponents.H,
			compositionFuelComponents.C,
			compositionFuelComponents.S,
			compositionFuelComponents.O,
			compositionFuelComponents.N,
			compositionFuelComponents.W,
			compositionFuelComponents.A,
		},
		DryComposition:        dryCompositionFuelComponents,
		BurnComposition:       burnCompositionFuelComponents,
		CoeffWorkDry: 		   coeffWorkDry,
		CoeffWorkBurn:		   coeffWorkBurn,
		HeatCombustion:        heatCombustion,
		DryFuelHeatCombustion: dryFuelHeatCombustion,
		BurnFuelHeatCombustion: burnFuelHeatCombustion,
	}
	// Display the result
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
	http.ListenAndServe(":8080", nil)
}

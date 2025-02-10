package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"lab1/logic"
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

	// Display the result
	fmt.Fprintf(w, "Dry Composition: %v\n", dryCompositionFuelComponents)
	fmt.Fprintf(w, "Burn Composition: %v\n", burnCompositionFuelComponents)

	tmpl.Execute(w, struct{ Success bool }{true})

	fmt.Fprint(w, "Hello world!")
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

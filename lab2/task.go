package main

import (
	"html/template"
	"lab1/logic"
	"log"
	"net/http"
	"strconv"
)

package main

import (
	"errors"
	"fmt"
)

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("task.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	var bCoal, bFuel, bGas float64
	fmt.Print("Введіть кількість вугілля (ГДж): ")
	_, err1 := fmt.Scan(&bCoal)
	fmt.Print("Введіть кількість мазуту (ГДж): ")
	_, err2 := fmt.Scan(&bFuel)
	fmt.Print("Введіть кількість газу (ГДж): ")
	_, err3 := fmt.Scan(&bGas)

	if errors.Is(err1, nil) && errors.Is(err2, nil) && errors.Is(err3, nil) {
		qCoal, gVynCoal, aVynCoal, aRCoal, nZuCoal, kHSCoal := 20.47, 1.5, 0.8, 25.2, 0.985, 0.0
		qFuel, gVynFuel, aVynFuel, aRFuel, nZuFuel, kHSFuel := 40.40, 0.0, 1.0, 0.15, 0.985, 0.0
		qGas := 20.47
		kGasValue := 0.0

		kCoalValue := logic.CalculateK(qCoal, aVynCoal, aRCoal, gVynCoal, nZuCoal, kHSCoal)
		eCoalValue := logic.CalculateE(kCoalValue, bCoal, qCoal)
		kFuelValue := logic.CalculateK(qFuel, aVynFuel, aRFuel, gVynFuel, nZuFuel, kHSFuel)
		eFuelValue := logic.CalculateE(kFuelValue, bFuel, qFuel)
		eGasValue := logic.CalculateE(kGasValue, bGas, qGas)

		fmt.Printf("Показник емісії твердих частинок (вугілля): %.2f г/ГДж\n", round(kCoalValue))
		fmt.Printf("Валовий викид (вугілля): %.2f т.\n", round(eCoalValue))
		fmt.Printf("Показник емісії твердих частинок (мазут): %.2f г/ГДж\n", round(kFuelValue))
		fmt.Printf("Валовий викид (мазут): %.2f т.\n", round(eFuelValue))
		fmt.Printf("Показник емісії твердих частинок (газ): %.2f г/ГДж\n", round(kGasValue))
		fmt.Printf("Валовий викид (газ): %.2f т.\n", round(eGasValue))
	} else {
		fmt.Println("Помилка: Усі введені значення мають бути числами.")
	}

	data := struct {
		Success bool
	}{
		Success: true,
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

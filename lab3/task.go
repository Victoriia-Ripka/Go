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

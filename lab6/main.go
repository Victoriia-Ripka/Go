package main

import (
	"lab6/pages/entry"
	"lab6/service"
	"net/http"
	"fmt"
)

var calcService = &service.CalculatorService{}

func main() {
	http.HandleFunc("/assets/testInputs.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "assets/testInputs.json")
	})
	http.HandleFunc("/assets/inputs.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "assets/testInputs.json")
	})
	http.HandleFunc("/assets/extraEPdata.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "assets/testInputs.json")
	})

	http.HandleFunc("/", entry.MainScreen)
	http.HandleFunc("/calculate", entry.CalculateHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

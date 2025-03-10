package main

import (
	"lab5/screens/calculators"
	"lab5/screens/entry"
	"lab5/services"
	"net/http"
	"fmt"
)

var calcService = &service.CalculatorService{}

func main() {
	http.HandleFunc("/", entry.MainScreen)
	http.HandleFunc("/calculator1", calculators.Calculator1Screen)
	http.HandleFunc("/calculator2", calculators.Calculator2Screen)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

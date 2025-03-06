package main

import (
	"lab5/screens/calculators"
	"lab5/screens/entry"
	"lab5/services"

	"html/template"
	"net/http"
)

var calcService = &service.CalculatorService{}

func main() {
	// Handle the different endpoints
	http.HandleFunc("/", mainScreen)
	http.HandleFunc("/calculator1", calculator1Screen)
	http.HandleFunc("/calculator2", calculator2Screen)

	// Start the web server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

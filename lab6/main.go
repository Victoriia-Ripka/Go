package main

import (
	"lab6/pages/entry"
	"lab6/service"
	"net/http"
	"fmt"
)

var calcService = &service.CalculatorService{}

func main() {
	http.HandleFunc("/", entry.MainScreen)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

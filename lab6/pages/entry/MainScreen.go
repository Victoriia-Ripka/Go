package entry

import (
	// "fmt"
	"html/template"
	"net/http"
	// "path/filepath"
	"lab6/service"
	"strconv"
	"math"
)

var calcService = &service.CalculatorService{}

type CalculationResult struct {
	KV float64
	Pp float64
	Sp float64
	Ip float64
}

func MainScreen(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/entry/entry.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	// Helper function to extract and parse form values
	parseFormValue := func(field string) float64 {
		value, err := strconv.ParseFloat(r.FormValue(field), 64)
		if err != nil {
			http.Error(w, "Invalid "+field+" value", http.StatusBadRequest)
			return 0
		}
		return value
	}

	// Parsing form values
	count := parseFormValue("count")
	capacity := parseFormValue("capacity")
	coefUsage := parseFormValue("coefUsage")
	coeffReactPower := parseFormValue("coeffReactPower")
	voltage := parseFormValue("voltage")
	coeffPower := parseFormValue("coeffPower")
	coeffUsefulAct := parseFormValue("coeffUsefulAct")

	// Creating input array
	epInputs := []service.EPInput{
		{
			Count:           count,
			Capacity:        capacity,
			CoefUsage:       coefUsage,
			CoeffReactPower: coeffReactPower,
			Voltage:         voltage,
			CoeffPower:      coeffPower,
			CoeffUsefulAct:  coeffUsefulAct,
		},
	}

	// Define constants
	// allCount := 81.0
	allNPh := 2330.0
	allNPK := 752.0
	allNPKtg := 657.0
	allNP2 := 96399.0

	// Step 3 calculations
	NPhList := calcService.CalculateNPh(epInputs) // 80, 28, ...
	NPhSum := calcService.CalculateSumNPh(NPhList) // 456
	// I := calcService.CalculateI(epInputs) // 146, 51, ...
	// sumCount := calcService.CalculateSumCount(epInputs) // 16

	// Step 4 calculations
	KV := calcService.CalculateGroupUtilizationCoeff(epInputs) // 0.208
	nE := calcService.CalculateEfCount(NPhSum, epInputs) // 15
	Kp := 1.25
	Pp := calcService.CalculatePp(Kp, epInputs) // 118
	Qp := calcService.CalculateQp(nE, epInputs) // 107
	Sp := calcService.CalculateSp(Pp, Qp) // 160
	Ip := calcService.CalculateIp(Pp, epInputs[3].Voltage) // 313

	// Step 6 calculations for all
	allKV := allNPK / allNPh
	allNe := math.Pow(allNPh, 2.0) / allNP2
	allKp := 0.7
	allPp := allKp * allNPK
	allQp := allKp * allNPKtg
	allSp := math.Sqrt(math.Pow(allPp, 2.0) + math.Pow(allQp, 2.0))
	allIp := allPp / epInputs[3].Voltage

	// Result array with rounded values
	resultArray := []float64{
		round(KV, 2),
		float64(nE),
		round(Kp, 2),
		round(Pp, 2),
		round(Qp, 2),
		round(Sp, 2),
		round(Ip, 2),
		round(allKV, 2),
		round(allNe, 2),
		round(allKp, 2),
		round(allPp, 2),
		round(allQp, 2),
		round(allSp, 2),
		round(allIp, 2),
	}

	result := CalculationResult{
		KV:  resultArray[0],
		Pp:  resultArray[3],
		Sp:  resultArray[5],
		Ip:  resultArray[6],
	}
	
	tmpl, err := template.ParseFiles("pages/entry/entry.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{"Result": result})
}

func round(value float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Round(value*shift) / shift
}
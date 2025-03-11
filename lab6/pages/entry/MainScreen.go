package entry

import (
	"fmt"
	"html/template"
	"net/http"
	// "path/filepath"
	"lab6/service"
	"strconv"
	"math"
)

var calcService = &service.CalculatorService{}

type CalculationResult struct {
    KV              float64
    NPhSum          float64
    Pp              float64
    Sp              float64
    Ip              float64
    AllKV           float64
    AllNe           float64
    AllKp           float64
    AllPp           float64
    AllQp           float64
    AllSp           float64
    AllIp           float64
    CoefUsage       float64
    CoeffReactPower float64
    Voltage         float64
    ResultArray     []float64
}

func MainScreen(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/entry/entry.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func createEPInput(r *http.Request, index int) service.EPInput {
	count, _ := strconv.ParseFloat(r.FormValue(fmt.Sprintf("ep%d-count", index)), 64)
	capacity, _ := strconv.ParseFloat(r.FormValue(fmt.Sprintf("ep%d-capacity", index)), 64)
	coefUsage, _ := strconv.ParseFloat(r.FormValue(fmt.Sprintf("ep%d-coefUsage", index)), 64)
	coeffReactPower, _ := strconv.ParseFloat(r.FormValue(fmt.Sprintf("ep%d-coeffReactPower", index)), 64)
	voltage, _ := strconv.ParseFloat(r.FormValue(fmt.Sprintf("ep%d-voltage", index)), 64)
	coeffPower, _ := strconv.ParseFloat(r.FormValue(fmt.Sprintf("ep%d-coeffPower", index)), 64)
	coeffUsefulAct, _ := strconv.ParseFloat(r.FormValue(fmt.Sprintf("ep%d-coeffUsefulAct", index)), 64)

	return service.EPInput{
		Count:           count,
		Capacity:        capacity,
		CoefUsage:       coefUsage,
		CoeffReactPower: coeffReactPower,
		Voltage:         voltage,
		CoeffPower:      coeffPower,
		CoeffUsefulAct:  coeffUsefulAct,
	}
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	var epInputs []service.EPInput
	for i := 1; i <= 10; i++ {
		epInputs = append(epInputs, createEPInput(r, i))
	}

    // Step 3 calculations
    // fmt.Println("Value of epInputs:", epInputs)
    NPhList := calcService.CalculateNPh(epInputs)
    // fmt.Println("Value of NPhList:", NPhList)
    NPhSum := calcService.CalculateSumNPh(NPhList)
    // fmt.Println("Value of NPhSum:", NPhSum)

    // Step 4 calculations
    KV := calcService.CalculateGroupUtilizationCoeff(epInputs)
    nE := calcService.CalculateEfCount(NPhSum, epInputs)
    Kp := 1.25
    Pp := calcService.CalculatePp(Kp, epInputs)
    Qp := calcService.CalculateQp(nE, epInputs)
    Sp := calcService.CalculateSp(Pp, Qp)
    Ip := calcService.CalculateIp(Pp, epInputs[0].Voltage)

    // Step 6 calculations for all
    allNPh := 2330.0
    allNPK := 752.0
    allNPKtg := 657.0
    allNP2 := 96399.0

    allKV := allNPK / allNPh
    allNe := math.Pow(allNPh, 2.0) / allNP2
    allKp := 0.7
    allPp := allKp * allNPK
    allQp := allKp * allNPKtg
    allSp := math.Sqrt(math.Pow(allPp, 2.0) + math.Pow(allQp, 2.0))
    allIp := allPp / epInputs[0].Voltage

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

    // Prepare the calculation result
    result := CalculationResult{
		KV:              resultArray[0],
		NPhSum:          NPhSum,
		Pp:              resultArray[3],
		Sp:              resultArray[5],
		Ip:              resultArray[6],
		AllKV:           resultArray[7],
		AllNe:           resultArray[8],
		AllKp:           resultArray[9],
		AllPp:           resultArray[10],
		AllQp:           resultArray[11],
		AllSp:           resultArray[12],
		AllIp:           resultArray[13],
		CoefUsage:       epInputs[0].CoefUsage,
		CoeffReactPower: epInputs[0].CoeffReactPower,
		Voltage:         epInputs[0].Voltage,
		ResultArray:     resultArray,
	}

    // Check if result array is not empty and first element > 0
    if len(resultArray) > 0 && resultArray[0] > 0.0 {
		tmpl, err := template.ParseFiles("pages/entry/entry.html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}
	
		tmpl.Execute(w, struct {
			Success bool
			Result CalculationResult
		}{
			Success: true,
			Result:     result,
		})
	} else {
		http.Error(w, "Invalid calculation result", http.StatusBadRequest)
	}
}

func round(value float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Round(value*shift) / shift
}
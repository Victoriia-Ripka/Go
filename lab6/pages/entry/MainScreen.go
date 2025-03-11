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
    allNPh := 2330.0
    allNPK := 752.0
    allNPKtg := 657.0
    allNP2 := 96399.0

    // Step 3 calculations
    NPhList := calcService.CalculateNPh(epInputs)
    NPhSum := calcService.CalculateSumNPh(NPhList)

    // Step 4 calculations
    KV := calcService.CalculateGroupUtilizationCoeff(epInputs)
    nE := calcService.CalculateEfCount(NPhSum, epInputs)
    Kp := 1.25
    Pp := calcService.CalculatePp(Kp, epInputs)
    Qp := calcService.CalculateQp(nE, epInputs)
    Sp := calcService.CalculateSp(Pp, Qp)
    Ip := calcService.CalculateIp(Pp, epInputs[0].Voltage)

    // Step 6 calculations for all
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
		CoefUsage:       coefUsage,
		CoeffReactPower: coeffReactPower,
		Voltage:         voltage,
		ResultArray:     resultArray, // Make sure ResultArray is included if you want to use it in the template
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
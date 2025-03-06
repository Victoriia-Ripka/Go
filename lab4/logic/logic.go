package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
)

// CalculatorService struct
type CalculatorService struct {
	CableData []CableData
}

// CableData struct to represent cable JSON structure
type CableData struct {
	Conductor string            `json:"conductor"`
	Type      string            `json:"type"`
	Time      map[string]float64 `json:"time"`
}

// NewCalculatorService initializes a new instance of CalculatorService
func NewCalculatorService(dataPath string) (*CalculatorService, error) {
	file, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, err
	}

	var cables []CableData
	err = json.Unmarshal(file, &cables)
	if err != nil {
		return nil, err
	}

	return &CalculatorService{CableData: cables}, nil
}

// roundUpToNearest rounds a number up to the nearest base
func roundUpToNearest(number float64, base int) int {
	return int(math.Ceil(number/float64(base)) * float64(base))
}

// findCableByConductorAndType finds a cable based on conductor and type
func (cs *CalculatorService) findCableByConductorAndType(conductor, cableType string) (*CableData, error) {
	for _, cable := range cs.CableData {
		if cable.Conductor == conductor && cable.Type == cableType {
			return &cable, nil
		}
	}
	return nil, errors.New("no matching cable found")
}

// findJekValue finds the Jek value for a given time range
func (cs *CalculatorService) findJekValue(selectedCable *CableData, timeInput float64) (float64, error) {
	timeRanges := []string{"1000-3000", "3000-5000", "5000-10000"}
	for _, rangeStr := range timeRanges {
		var start, end int
		fmt.Sscanf(rangeStr, "%d-%d", &start, &end)
		if timeInput >= float64(start) && timeInput <= float64(end) {
			return selectedCable.Time[rangeStr], nil
		}
	}
	return 0, errors.New("no matching time range found")
}

// CalculateCablesCompatibility computes the required cable section
func (cs *CalculatorService) CalculateCablesCompatibility(conductor, cableType string, timeInput float64) (int, error) {
	_, IK, tF, _ := 10.0, 2.5*1000, 2.5, 1300.0

	selectedCable, err := cs.findCableByConductorAndType(conductor, cableType)
	if err != nil {
		return 0, err
	}

	_, err = cs.findJekValue(selectedCable, timeInput)
	if err != nil {
		return 0, err
	}

	// IM := (SM / 2) / (math.Sqrt(3) * U)
	Smin := IK * math.Sqrt(tF) / 92

	return roundUpToNearest(Smin, 10), nil
}

// DeterminateCurrent calculates the short-circuit current
func (cs *CalculatorService) DeterminateCurrent(u, sk, sNomt float64) float64 {
	Xc := math.Pow(u, 2) / sk
	Xt := (u / 100) * (math.Pow(u, 2) / sNomt)
	X := Xc + Xt
	Ip0 := u / (math.Sqrt(3) * X)
	return math.Round(Ip0*100) / 100
}

// DeterminateSubstationCurrent computes various substation currents
func DeterminateSubstationCurrent() [4]float64 {
	uVn, sNomt, uKMax := 115.0, 6.3, 11.1
	Rsn110, Xcn110, Rsmin110, Xcmin110 := 10.65, 24.02, 34.88, 65.68

	xT := (uKMax * math.Pow(uVn, 2)) / (100 * sNomt)
	rSH, xSH := Rsn110, Xcn110+xT
	zSH := math.Hypot(rSH, xSH)

	rSHmin, xSHmin := Rsmin110, Xcmin110+xT
	zSHmin := math.Hypot(rSHmin, xSHmin)

	i3SH := uVn * 1000 / (math.Sqrt(3) * zSH)
	i2SH := i3SH * math.Sqrt(3) / 2
	i3SHmin := uVn * 1000 / (math.Sqrt(3) * zSHmin)
	i2SHmin := i3SHmin * math.Sqrt(3) / 2

	return [4]float64{
		math.Round(i3SH*100) / 100,
		math.Round(i2SH*100) / 100,
		math.Round(i3SHmin*100) / 100,
		math.Round(i2SHmin*100) / 100,
	}
}
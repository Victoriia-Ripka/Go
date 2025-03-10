package service

import (
	"fmt"
)

type CalculatorService struct{}

func (cs *CalculatorService) CompareReliabilitySystems() [2]float64 {
	wSv := 0.02
	w1 := [10]float64{0.01, 0.07, 0.015, 0.02, 0.03, 0.03, 0.03, 0.03, 0.03, 0.03}
	tV1 := [10]float64{30, 10, 100, 15, 2, 2, 2, 2, 2, 2}
	wOs := sum(w1[:])
	tVos := weightedAverage(w1[:], tV1[:], wOs)

	kAos := wOs * tVos / 8760
	kPos := 1.2 * 43 / 8760

	wDk := 2 * wOs * (kAos + kPos)
	wDs := wDk + wSv

	roundedValue1 := round(wDk, 4)
	roundedValue2 := round(wDs, 4)

	// fmt.Printf("Calculator1 - w_dk: %.4f\n", wDk)
	// fmt.Printf("Calculator1 - w_ds: %.4f\n", wDs)

	return [2]float64{roundedValue1, roundedValue2}
}

func (cs *CalculatorService) CalculateLosses() string {
	zPera := 23.6
	zPerp := 17.6
	w := 0.01
	tV := 0.045

	tM := 5.12 * 1000 * 6451
	kP := 0.004

	mWNeda := w * tV * tM
	mWNedp := kP * tM

	mZPer := zPera*mWNeda + zPerp*mWNedp
	return fmt.Sprintf("%.2f", mZPer)
}

func sum(arr []float64) float64 {
	var total float64
	for _, value := range arr {
		total += value
	}
	return total
}

func weightedAverage(weights, times []float64, weightSum float64) float64 {
	var total float64
	for i := range weights {
		total += weights[i] * times[i]
	}
	return total / weightSum
}

func round(value float64, places int) float64 {
	format := fmt.Sprintf("%%.%df", places)
	roundedValue := fmt.Sprintf(format, value)
	var rounded float64
	fmt.Sscanf(roundedValue, "%f", &rounded)
	return rounded
}
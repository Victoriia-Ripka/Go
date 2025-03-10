package service

import (
	"fmt"
)

type CalculatorService struct{}

// CompareReliabilitySystems performs the calculations for the reliability systems and returns an array of two values.
func (cs *CalculatorService) CompareReliabilitySystems() [2]float64 {
	// Initial values
	// nSwitcher1 := 110
	// l1 := 10
	// transformerN1 := [2]int{110, 10}
	// inputSwitch1 := 10
	// accesionN1 := 10
	// accesionCount1 := 6

	// nSwitcher2 := 110
	// l2 := 10
	// transformerN2 := [2]int{110, 10}
	// inputSwitch2 := 10
	// accesionN2 := 10
	// accesionCount2 := 6
	// sectionalSwitchN2 := 10

	wSv := 0.02

	// Weights and time values
	w1 := [10]float64{0.01, 0.07, 0.015, 0.02, 0.03, 0.03, 0.03, 0.03, 0.03, 0.03}
	tV1 := [10]float64{30, 10, 100, 15, 2, 2, 2, 2, 2, 2}

	// Sum of weights and weighted average of time values
	wOs := sum(w1[:])
	tVos := weightedAverage(w1[:], tV1[:], wOs)

	kAos := wOs * tVos / 8760
	kPos := 1.2 * 43 / 8760

	wDk := 2 * wOs * (kAos + kPos)
	wDs := wDk + wSv

	// Return the rounded values
	roundedValue1 := round(wDk, 4)
	roundedValue2 := round(wDs, 4)

	// Print results (equivalent to Log.d in Android)
	fmt.Printf("Calculator1 - w_dk: %.4f\n", wDk)
	fmt.Printf("Calculator1 - w_ds: %.4f\n", wDs)

	return [2]float64{roundedValue1, roundedValue2}
}

// CalculateLosses calculates the losses and returns the result as a string.
func (cs *CalculatorService) CalculateLosses() string {
	// Constants
	zPera := 23.6
	zPerp := 17.6
	w := 0.01
	tV := 0.045

	// Other constants
	tM := 5.12 * 1000 * 6451
	kP := 0.004

	// Loss calculations
	mWNeda := w * tV * tM
	mWNedp := kP * tM

	// Final result
	mZPer := zPera*mWNeda + zPerp*mWNedp
	return fmt.Sprintf("%.2f", mZPer)
}

// Helper function to calculate sum of an array
func sum(arr []float64) float64 {
	var total float64
	for _, value := range arr {
		total += value
	}
	return total
}

// Helper function to calculate weighted average
func weightedAverage(weights, times []float64, weightSum float64) float64 {
	var total float64
	for i := range weights {
		total += weights[i] * times[i]
	}
	return total / weightSum
}

// Helper function to round a float to the specified number of decimal places

func round(value float64, places int) float64 {
	// Format the float to a string with the specified number of decimal places
	format := fmt.Sprintf("%%.%df", places)
	roundedValue := fmt.Sprintf(format, value)
	var rounded float64
	// Convert the formatted string back to a float64
	fmt.Sscanf(roundedValue, "%f", &rounded)
	return rounded
}
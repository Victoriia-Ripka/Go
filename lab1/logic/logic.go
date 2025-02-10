package logic

import "math"

func CalculateCoeffWorkDry(w float64) float64 {
	return 100.0 / (100.0 - w)
}

func CalculateCoeffDryWork(w float64) float64 {
	return (100.0 - w) / 100.0
}

func CalculateCoeffWorkBurn(w, a float64) float64 {
	return 100.0 / (100.0 - w - a)
}

func CalculateCoeffBurnWork(w, a float64) float64 {
	return (100.0 - w - a) / 100.0
}

func CalculateNewPercentage(coeff float64, elements []float64) []float64 {
	newElPercentage := make([]float64, len(elements))
	for index, value := range elements {
		newElPercentage[index] = math.Round(value*coeff*100) / 100
	}

	return newElPercentage
}

func CalculateHeatCombustion(C, H, O, S, W float64) float64 {
	return math.Round((339 * C + 1030 * H - 108.8 * ( O - S ) - 25 * W)*100) / 100
}

func CalculateDryFuelHeatCombustion(heatCombustion, W, workDryCoeff float64) float64 {
	return math.Round((heatCombustion + 0.025 * W) * workDryCoeff*100) / 100
}

func CalculateBurnFuelHeatCombustion(heatCombustion, W, workBurnCoeff float64) float64 {
	return math.Round((heatCombustion + 0.025 * W) * workBurnCoeff*100) / 100
}

func CalculateWorkFuelPercentage(burnWorkCoeff, dryWorkCoeff, C, H, O, S, A, V float64) []float64 {
	workCarbon := C * burnWorkCoeff
	workHydrogen := H * burnWorkCoeff
	workOxygen := O * burnWorkCoeff
	workSulfur := S * burnWorkCoeff
	workAsh := A * dryWorkCoeff
	workVanadium := V * dryWorkCoeff

	return []float64{workCarbon, workHydrogen, workOxygen, workSulfur, workAsh, workVanadium}
}
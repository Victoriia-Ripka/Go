package logic

func CalculateCoeffWorkDry(w float64) float64 {
	return 100.0 / (100.0 - w)
}

func CalculateCoeffWorkBurn(w, a float64) float64 {
	return 100.0 / (100.0 - w - a)
}

func CalculateNewPercentage(coeff float64, elements []float64) []float64 {
	newElPercentage := make([]float64, len(elements))
	for index, value := range elements {
		newElPercentage[index] = value * coeff
	}

	return newElPercentage
}

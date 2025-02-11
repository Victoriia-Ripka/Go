package logic

import "math"

func CalculateW(p, sigmaWValue float64, mode string) float64 {
	if mode == "with" {
		return p * 24 * sigmaWValue
	} else if mode == "without" {
		return p * 24 * (1 - sigmaWValue)
	}
	return 0.0
}

func ProfFineCalculate(w, b float64) float64 {
	return w * b
}

func TrapezoidalRule(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		sum += f(x)
	}
	return h * sum
}

func CalculateP(PcValue, sigmaValue, BValue float64) float64 {
	x, y := PcValue*0.95, PcValue*1.05

	pdFunctionUpdated := func(p float64) float64 {
		return math.Exp(-math.Pow(p-PcValue, 2)/(0.25*0.25*2)) / (0.25 * math.Sqrt(2*math.Pi))
	}
	sigmaWValueUpdated := TrapezoidalRule(pdFunctionUpdated, x, y, 100)

	w3 := CalculateW(PcValue, sigmaWValueUpdated, "with")
	prof2 := ProfFineCalculate(w3, BValue)
	w4 := CalculateW(PcValue, sigmaWValueUpdated, "without")
	fine2 := ProfFineCalculate(w4, BValue)

	profitValue := prof2 - fine2
	return Round(profitValue)
}

func Round(value float64) float64 {
	return math.Round(value*100) / 100
}
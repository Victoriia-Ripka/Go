package logic

import "math"

func CalculateK(q, aVyn, aR, gVyn, nZu, kHS float64) float64 {
	return (math.Pow(10, 6) * aVyn * aR) / (q * (100 - gVyn)) * (1 - nZu) + kHS
}

func CalculateE(k, b, q float64) float64 {
	return math.Pow(10, -6) * k * b * q
}

func Round(value float64) float64 {
	return math.Round(value*100) / 100
}
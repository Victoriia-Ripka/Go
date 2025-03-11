package service

import (
	"math"
)

type EPInput struct {
	Count           float64
	Capacity        float64
	CoefUsage       float64
	CoeffReactPower float64
	Voltage         float64
	CoeffPower      float64
	CoeffUsefulAct  float64
}

type CalculatorService struct{}

func (cs *CalculatorService) calculateNPK(epInputs []EPInput) []float64 {
	nPK := make([]float64, len(epInputs))
	for i, input := range epInputs {
		nPK[i] = input.Count * input.Capacity * input.CoefUsage
	}
	return nPK
}

func (cs *CalculatorService) calculatePKTgSum(epInputs []EPInput) float64 {
	nPK := cs.calculateNPK(epInputs)
	var sum float64
	for i, input := range epInputs {
		sum += input.CoeffReactPower * nPK[i]
	}
	return sum
}

func (cs *CalculatorService) CalculateNPh(epInputs []EPInput) []float64 {
	NPhList := make([]float64, len(epInputs))
	for i, input := range epInputs {
		NPhList[i] = input.Count * input.Capacity
	}
	return NPhList
}

func (cs *CalculatorService) CalculateSumNPh(NPhList []float64) float64 {
	return sum(NPhList)
}

func (cs *CalculatorService) CalculateI(epInputs []EPInput) []float64 {
	I := make([]float64, len(epInputs))
	for i, input := range epInputs {
		if input.Voltage == 0 || input.CoeffPower == 0 || input.CoeffUsefulAct == 0 {
			I[i] = 0
		} else {
			NPh := input.Count * input.Capacity
			I[i] = NPh / (math.Sqrt(3) * input.Voltage * input.CoeffPower * input.CoeffUsefulAct)
		}
	}
	return I
}

func (cs *CalculatorService) CalculateSumCount(epInputs []EPInput) int {
	var sum int
	for _, input := range epInputs {
		sum += int(input.Count)
	}
	return sum
}

func (cs *CalculatorService) CalculateGroupUtilizationCoeff(epInputs []EPInput) float64 {
	nPK := cs.calculateNPK(epInputs)
	var nPsum float64
	for _, input := range epInputs {
		nPsum += input.Count * input.Capacity
	}
	return sum(nPK) / nPsum
}

func (cs *CalculatorService) CalculateEfCount(NPh float64, epInputs []EPInput) int {
	var sum float64
	for _, input := range epInputs {
		sum += input.Count * math.Pow(input.Capacity, 2)
	}
	return int(math.Round(math.Pow(NPh, 2)/sum)) + 1
}

func (cs *CalculatorService) CalculatePp(Kp float64, epInputs []EPInput) float64 {
	nPK := cs.calculateNPK(epInputs)
	return Kp * sum(nPK)
}

func (cs *CalculatorService) CalculateQp(en int, epInputs []EPInput) float64 {
	kptg := cs.calculatePKTgSum(epInputs)
	if en <= 10 {
		return kptg * 1.1
	}
	return kptg
}

func (cs *CalculatorService) CalculateSp(Pp, Qp float64) float64 {
	return math.Sqrt(math.Pow(Pp, 2) + math.Pow(Qp, 2))
}

func (cs *CalculatorService) CalculateIp(Pp, Up float64) float64 {
	if Up == 0 {
		return 0
	}
	return Pp / Up
}

// func parseToInt(value string) int {
// 	var result int
// 	fmt.Sscanf(value, "%d", &result)
// 	return result
// }

// func parseToFloat(value string) float64 {
// 	var result float64
// 	fmt.Sscanf(value, "%f", &result)
// 	return result
// }

func sum(arr []float64) float64 {
	var result float64
	for _, v := range arr {
		result += v
	}
	return result
}

package utils

import "math"

func GetInterestRatePerTerm(annualInterestRate float64) float64 {
	return annualInterestRate/12
}

func GetDueDateAmount(totalMoney int, interestRate float64, terms int) int {
	calc := float64(totalMoney) * (interestRate / (1 - math.Pow(1 + interestRate, -float64(terms))))
	if terms == 1 {
		return int(math.Ceil(calc))
	}
	return int(math.Floor(calc))
}

func GetInterestAmountByTerm(remainingMoney int, interestRate float64) int {
	return int(math.Round(float64(remainingMoney) * interestRate))
}
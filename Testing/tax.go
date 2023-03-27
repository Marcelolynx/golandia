package testing

func CalculateTax(amount float64) float64 {
	if amount >= 1000 {
		return amount * 0.2
	}
	return amount * 0.1
}

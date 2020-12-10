package calculations

func HornerScheme(coefficients []float64, x float64) float64 {
	total := coefficients[len(coefficients)-1]
	for i := len(coefficients) - 2; i >= 0; i-- {
		total = total*x + coefficients[i]
	}
	return total
}

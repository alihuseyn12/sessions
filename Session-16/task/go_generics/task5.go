package go_generics

func Sum[T int | float64](slice []T) T {
	var total T
	for _, v := range slice {

		total += v
	}
	return total
}

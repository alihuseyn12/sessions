package go_generics

func MinMax[T int | float64](slice []T) (T, T) {
	var min T = slice[0]
	var max T = slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v

		}
	}
	return min, max

}

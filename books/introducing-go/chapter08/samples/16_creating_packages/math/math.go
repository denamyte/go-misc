package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Min(xs []float64) float64 {
	min := xs[0]
	for _, x := range xs[1:] {
		if x < min {
			min = x
		}
	}
	return min
}

func Max(xs []float64) float64 {
	max := xs[0]
	for _, x := range xs[1:] {
		if x > max {
			max = x
		}
	}
	return max
}

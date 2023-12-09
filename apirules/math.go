package apirules

//// Max finds the maximum value in a slice of float64 values. It initializes a variable maxVal with negative infinity and iterates over the slice, comparing each value with maxVal. If
//func (inst *RQL) Max(s ...float64) float64 {
//	maxVal := math.Inf(-1)
//	for _, v := range s {
//		if v > maxVal {
//			maxVal = v
//		}
//	}
//	return maxVal
//}
//
//// Min finds the minimum value in a slice of float64 values. It initializes minVal as positive infinite and then iterates through the values in the slice, updating minVal if a smaller
//func (inst *RQL) Min(s ...float64) float64 {
//	minVal := math.Inf(1)
//	for _, v := range s {
//		if v < minVal {
//			minVal = v
//		}
//	}
//	return minVal
//}
//
//// Avg calculates the average value for the given slice of float64 values.
//func (inst *RQL) Avg(s ...float64) float64 {
//	total := 0.0
//	for _, v := range s {
//		total += v
//	}
//	return total / float64(len(s))
//}

// max finds the maximum value in a slice of float64 numbers.
func max(s []float64) float64 {
	maxVal := s[0]
	for _, v := range s {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}

// min finds the minimum value in a slice
func min(s []float64) float64 {
	minVal := s[0]
	for _, v := range s {
		if v < minVal {
			minVal = v
		}
	}

	return minVal
}

// Finds the average value of a slice
func avg(s []float64) float64 {
	total := 0.0
	for _, v := range s {
		total += v
	}

	return total / float64(len(s))
}

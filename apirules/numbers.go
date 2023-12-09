package apirules

import (
	"math"
	"math/rand"
	"time"
)

// RandInt returns a random int within the specified range.
func (inst *RQL) RandInt(range1, range2 int) int {
	if range1 == range2 {
		return range1
	}
	var min, max int
	if range1 > range2 {
		max = range1
		min = range2
	} else {
		max = range2
		min = range1
	}
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RandFloat returns a random float64 within the specified range.
func (inst *RQL) RandFloat(range1, range2 float64) float64 {
	if range1 == range2 {
		return range1
	}
	var min, max float64
	if range1 > range2 {
		max = range1
		min = range2
	} else {
		max = range2
		min = range1
	}
	rand.Seed(time.Now().UnixNano())
	return inst.RoundTo(min+rand.Float64()*(max-min), 2)
}

// LimitToRange returns the input value restricted to the specified range.
// If the range is a single value, that value is returned.
// The minimum and maximum values of the range are determined based on range1 and range2 parameters.
// If range1 is greater than range2, range1 is considered the maximum and range2 is considered the minimum.
// If range2 is greater than range1, range2 is considered the maximum and range1 is considered the minimum.
// The input value is checked against the minimum and maximum values of the range using math.Min and math.Max functions.
// The result is the input value if it is within the range, otherwise the minimum or maximum value of the range is returned.
func (inst *RQL) LimitToRange(value float64, range1 float64, range2 float64) float64 {
	if range1 == range2 {
		return range1
	}
	var min, max float64
	if range1 > range2 {
		max = range1
		min = range2
	} else {
		max = range2
		min = range1
	}
	return math.Min(math.Max(value, min), max)
}

// RoundTo rounds the value to the specified number of decimal places.
// It takes a float64 value and a uint32 indicating the number of decimal places to round to.
// If the decimals parameter is less than 0, the original value is returned without rounding.
// The function uses the math.Round function to perform the rounding.
// The value is multiplied by 10 raised to the power of the decimal places, rounded, and then divided
// by 10 raised to the power of the decimal places to obtain the rounded value.
func (inst *RQL) RoundTo(value float64, decimals uint32) float64 {
	if decimals < 0 {
		return value
	}
	return math.Round(value*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

// Scale scales the given value from the input range to the output range.
// If either the input range or the output range has no variation (inMin == inMax or outMin == outMax),
// the value is returned unchanged.
// The scaled value is calculated as ((value - inMin) / (inMax - inMin)) * (outMax - outMin) + outMin.
// If the scaled value exceeds the maximum of the output range, the maximum value is returned.
// If the scaled value is smaller than the minimum of the output range, the minimum value is returned.
// Otherwise, the scaled value is returned.
func (inst *RQL) Scale(value, inMin, inMax, outMin, outMax float64) float64 {
	if inMin == inMax || outMin == outMax {
		return value
	}
	scaled := ((value-inMin)/(inMax-inMin))*(outMax-outMin) + outMin
	if scaled > math.Max(outMin, outMax) {
		return math.Max(outMin, outMax)
	} else if scaled < math.Min(outMin, outMax) {
		return math.Min(outMin, outMax)
	} else {
		return scaled
	}
}

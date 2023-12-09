package apirules

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want float64
	}{
		{
			name: "Single Value",
			s:    []float64{1},
			want: 1,
		},
		{
			name: "All Positive Values",
			s:    []float64{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "All Negative Values",
			s:    []float64{-1, -2, -3, -4, -5},
			want: -1,
		},
		{
			name: "Zero and Positive Values",
			s:    []float64{0, 1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "Zero and Negative Values",
			s:    []float64{0, -1, -2, -3, -4, -5},
			want: 0,
		},
		{
			name: "Floating Point Values",
			s:    []float64{-1.2, -2.3, -3.4, -4.5, -5.6},
			want: -1.2,
		},
		{
			name: " Decimal and Fractional Values",
			s:    []float64{0.5, 1.1, 2.2, 3.3, 4.4},
			want: 4.4,
		},
		{
			name: "Values with Max float64",
			s:    []float64{0, 1, 2, 3, 4, math.MaxFloat64},
			want: math.MaxFloat64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.s); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

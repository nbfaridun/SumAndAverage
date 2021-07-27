package main

import "testing"

func TestGetAverageAndSum(t *testing.T) {
	for _, tt := range []struct {
		x      []float64
		Sum float64
		Average float64

	}{
		{x: []float64{1, 15, 25, 37}, Sum: 78,  Average: 19.5},
		{x: []float64{0, 9, 36, 100}, Sum: 145, Average: 36.25},
		{x: []float64{2, 4, 6}, Sum: 12, Average: 4},
		{x: []float64{100, 25, -52, 1}, Sum: 74, Average: 18.5},
		{x: []float64{1}, Sum: 1, Average: 1},
	} {
		if ans1, ans2 := getSumAndAverage(tt.x...); ans1 != tt.Sum && ans2 != tt.Average {
			t.Fatalf("expected sum and average of %v to be %0.2f and %0.2f, got %0.2f and %0.2f\n", tt.x, tt.Sum, tt.Average, ans1, ans2)
		}
	}
}
package main

import "fmt"

func getSumAndAverage(x ...float64) (float64, float64) {

	var sum float64
	sum = 0

	for _, sliceElement := range x {
		sum += sliceElement
	}

	if sum == 0 {
		return 0, 0
	} else {
		return sum, sum / float64(len(x))
	}
}

func main() {
	var x = []float64{32, 38, 29, 42, 52, 2}

	ans1, ans2 := getSumAndAverage(x...)

	fmt.Println(ans1, ans2)
}


package main

import "fmt"

/** 1. Напишите функцию что принимает массив и возвращает сумму массива **/
/** 2. Нахождение среднего значения данного слайса функцией **/

// К функциям написать ЮнитТесты на гитхабе

func getSum(a []float64, k int) float64 {
	var i int
	var sum float64

	sum, i = 0, 0

	for i < k {
		sum += a[i]
		i++
	}

	return sum
}

func getAverage(a []float64, k int) float64 {
	var i int
	var sum float64

	sum, i = 0, 0

	for i < k {
		sum += a[i]
		i++
	}

	return sum / float64(i)
}

func main() {
	var x = []float64{32, 38, 29, 42, 52, 2}

	ans1 := getSum(x, 6)
	ans2 := getAverage(x, 6)

	fmt.Println(ans1, ans2)
}


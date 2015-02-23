package main

import "fmt"

func main() {
	a := []float32{1.1, 1.2, 1.3}
	sum, avg := SumAndAverage(a)
	fmt.Printf("Sum: %f, Avg: %f\n", sum, avg)
}

func SumAndAverage(list []float32) (sum,avg float32) {
	for _, value := range list {
		sum += value
	}
	avg = sum / float32(len(list))
	return
}
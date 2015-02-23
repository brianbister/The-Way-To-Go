package main

import "fmt"

func main() {
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(10))
}

func fibonacci(n int) []int {
	fibSlice := make([]int, n)
	for i := range fibSlice {
		if i == 0 || i == 1 {
			fibSlice[i] = 1
		} else {
			fibSlice[i] = fibSlice[i-1] + fibSlice[i-2]
		}
	}
	return fibSlice
}
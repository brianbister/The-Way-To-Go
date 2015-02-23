package main

import "fmt"

func main() {
	var fib [50]int
	for i := range fib {
		if i == 0 || i == 1 {
			fib[i] = 1
		} else {
			fib[i] = fib[i-1] + fib[i-2]
		}
	}
	fmt.Println(fib)
}
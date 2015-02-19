package main

import "fmt"

var sum, product, diff int

func main() {
	sum, product, diff = unnamed(2, 3)
	fmt.Printf("%d %d %d\n", sum, product, diff)
	sum, product, diff = named(2, 3)
	fmt.Printf("%d %d %d\n", sum, product, diff)
}

func unnamed(a int, b int) (int, int, int) {
	return a + b, a * b, b - a
}

func named(a int, b int) (sum int, prod int, diff int) {
	sum = a + b
	prod = a * b
	diff = b - a
	return
}
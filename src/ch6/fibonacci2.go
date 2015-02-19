package main

import "fmt"

func main() {
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(10))
}

func fibonacci(n int) (position int, value int) {
	position = n
	if n <= 1 {
		value = 1
		return
	}
	_, value1 := fibonacci(n-1)
	_, value2 := fibonacci(n-2)
	value = value1 + value2
	return
}
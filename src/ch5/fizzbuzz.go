package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Printf("FIZZBUZZ")
		case i % 3 == 0:
			fmt.Printf("FIZZ")
		case i % 5 == 0:
			fmt.Printf("BUZZ")
		default:
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
}
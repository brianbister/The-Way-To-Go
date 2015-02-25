package main

import (
	"fmt"
	"./even"
)

func main() {
	for i:=1; i<=100; i++ {
		if even.IsEven(i) {
			fmt.Printf("%d is even.\n", i)
		} else {
			fmt.Printf("%d is odd.\n", i)
		}
	}
}
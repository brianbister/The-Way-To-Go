package main

import "fmt"

func main() {
	PrintInts(1, 2, 3)
}

func PrintInts(a ...int) {
	for _, num := range a {
		fmt.Println(num)
	}
}
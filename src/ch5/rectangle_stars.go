package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		fmt.Printf("*")
	}
	fmt.Printf("\n")
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			if j == 0 || j == 19 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	for i := 0; i < 20; i++ {
		fmt.Printf("*")
	}
	fmt.Printf("\n")
}
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 2
	fmt.Println(a) // [1, 2, 3]
	fmt.Println(b) // [2, 2, 3]
}
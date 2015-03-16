package main

import "fmt"

type AnonStruct struct {
	floatField float64
	string
	int
}

func main() {
	a := AnonStruct{floatField: 1.5, string: "example", int:1}
	fmt.Println(a)
}
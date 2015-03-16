package main

import (
	"fmt"
	"strconv"
)

type Celsius float64

func (celsius Celsius) String() string {
	return strconv.FormatFloat(float64(celsius), 'f', 1, 64) + " °C"
}

func main() {
	var c Celsius = 15.2
	fmt.Println(c)
}
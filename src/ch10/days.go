package main

import (
	"fmt"
)

type Day int

const (
    MO Day = iota
    TU
    WE
    TH
    FR
    SA
    SU
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func main() {
	var day Day = MO
	fmt.Println(day)
}
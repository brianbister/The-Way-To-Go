package main

import "fmt"

type Employee struct {
	salary float64
}

func (e *Employee) giveRaise(percent float64) {
	e.salary += e.salary *  percent
}

func main() {
	employee := &Employee{salary: 100.0}
	fmt.Println(employee) // 100
	employee.giveRaise(1)
	fmt.Println(employee) // 200
}
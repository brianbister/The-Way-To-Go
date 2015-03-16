package main

import "fmt"

type Base struct {
	id int
}

func (base *Base) Id() int {
	return base.id
}

func (base *Base) SetId(id int) {
	base.id = id
}

type Person struct {
	firstName string
	lastName string
	Base
}

type Employee struct {
	salary float64
	Person
}

func main() {
	employee := new(Employee)
	employee.SetId(10)
	fmt.Println(employee.Id())
}
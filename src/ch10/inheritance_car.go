package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

type Mercedes struct {
	Car
}

func (car *Car) numberOfWheels() int {
	return car.wheelCount
}

func (car *Car) Start() {
	fmt.Println("The car has started.")
}

func (car *Car) Stop() {
	fmt.Println("The car has stopped.")
}

func (car *Car) GoToWorkIn() {
	car.Start()
	car.Stop()
}

func (mercedes *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Merkel!")
}

func main() {
	mercedes := &Mercedes{Car{wheelCount: 10}}
	fmt.Println(mercedes.numberOfWheels())
	mercedes.GoToWorkIn()
	mercedes.sayHiToMerkel()
}
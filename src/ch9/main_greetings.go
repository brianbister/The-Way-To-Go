package main

import (
	"fmt"
	"./greetings"
)

func main() {
	if greetings.IsAM() {
		fmt.Println("Good Morning")
	} else {
		fmt.Println("Good Evening")
	}
}
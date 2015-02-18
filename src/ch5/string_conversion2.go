package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "3"

	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("An error occurred, %s\n", err)
		return
	}
	fmt.Printf("The integer is %d\n", an)
}
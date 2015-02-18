package main

import "fmt"

func main() {
	gs := "G"
	for i := 0; i < 25; i, gs = i+1, gs+"G" {
		fmt.Println(gs)
	}
}
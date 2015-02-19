package main

import "fmt"

func main() {
	printrec(10)
}

func printrec(i int) {
	if i < 1 {
		return
	}
	fmt.Println(i)
	printrec(i-1)
}
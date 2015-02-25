package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	fmt.Printf("Size of Int in Bytes: %d\n", unsafe.Sizeof(a))
}
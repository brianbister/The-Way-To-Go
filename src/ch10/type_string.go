package main

import (
	"fmt"
)

type T struct {
	a int
	b float32
	c string
}

func (t T) String() string{
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}

func main() {
	t := &T{a:7, b:-2.35, c: "abc\tdef"}
	fmt.Println(t)
}
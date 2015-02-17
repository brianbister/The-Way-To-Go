package main

import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println(count_ascii("Hello World!")) // 12 bytes
	fmt.Println(count_utf8("Hello World!")) // 12 runes
	fmt.Println(count_ascii("你好")) // 6 bytes
	fmt.Println(count_utf8("你好")) // 2 runes
}

func count_ascii(s string) int {
	return len(s)
}

func count_utf8(s string) int {
	return utf8.RuneCountInString(s)
}
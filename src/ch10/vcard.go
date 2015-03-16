package main

import (
	"fmt"
	"time"
)

type Address struct {
	street string
	city string
}

type VCard struct {
	name string
	birthDate time.Time
	photo string
	addresses map[string]*Address
}

func main() {
	vcard := new(VCard)
	vcard.name = "Brian"
	vcard.photo = "photo.jpg"
	vcard.birthDate = time.Date(1991, 2, 26, 0, 0, 0, 0, time.Local)
	vcard.addresses = make(map[string]*Address)
	vcard.addresses["current"] = &Address{street:"6th", city:"Los Angeles"}
	fmt.Println(vcard)
}
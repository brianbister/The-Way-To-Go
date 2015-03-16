package main

type Person struct {
	firstName string
	lastName string
}

func upPerson (p Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {

}
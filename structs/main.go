package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zip int
}

func main() {
	 jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zip: 11020,
		},
	 }
	 //jimPointer := &jim
	 //jimPointer.updateName("Alex")
	 
	 jim.updateName("Alex")
	 jim.print()
}

func (p *person) updateName(newFirstName string){
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
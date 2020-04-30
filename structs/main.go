package main

import "fmt"

// Defining a struct a person struct here
// STRUCT is similar to obj in JS

type person struct {
	firstName   string
	lastName    string
	contactInfo contact
}

type contact struct {
	email   string
	zipcode int
}

func main() {

	// Ways to define struct
	// person1 := person{"Manoj", "Reddy"}
	// person2 := person{firstName: "Sirisha", lastName: "Reddy"}
	// fmt.Println("Persons are", person1.firstName, person2.firstName)

	var person3 person

	person3.firstName = "Manoj Reddy"
	person3.lastName = "Gangavarapu"
	person3.contactInfo.zipcode = 524223
	person3.contactInfo.email = "manoj.g@gmail.com"

	// person4 := person{
	// 	firstName: "test4",
	// 	lastName:  "testlast",
	// 	contactInfo: contact{
	// 		email:   "email",
	// 		zipcode: 333,
	// 	},
	// }

	// To get all the values of a struct in a nice way
	// fmt.Printf("%+v", person3)

	// fmt.Printf("%+v", person4)

	// person3.print()
	person3FullName := person3.fullName()
	fmt.Println(person3FullName)

	// First get the pointer
	// &variable -> gives us the memory address of the variable
	// person3Pointer := &person3
	// person3Pointer.updateName("Kumar")
	// the above two lines can also be called as
	// Go allows to call updateName either on type or pointer of that type
	// because we mentioned reciver as *person
	person3.updateName("Kumar")
	fmt.Println(person3)

}

// Way to add methods to a struct

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}

// How to update a struct basically modifying a struct
// POINTERS

// *person indiactes that we are working with a pointer to the person
// and also indiactes that updateName can be called only with a pointer
// simply means updateName can be only called with a person pointer
func (p *person) updateName(firstName string) {
	// *pointer gives us the values of the variable at this memory address
	(*p).firstName = firstName
}

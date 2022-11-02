package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// john := person{"John", "Doe"}
	// jane := person{firstName: "Jane", lastName: "Doe"}

	// Print whole property
	// fmt.Println(john, jane)

	// Print firstname of john and jane var
	//fmt.Println(john.firstName, jane.firstName)

	// Update value of struct
	// john.firstName = "Oskar"
	// fmt.Println(john.firstName, jane.firstName)

	person1 := person{
		firstName: "Oskar",
		lastName:  "Wulff",
		contactInfo: contactInfo{
			email:   "oskar9899@gmail.com",
			zipCode: 19547,
		},
	}

	person1.updateName("Updated value")
	person1.print()

	mySlice := []string{"hello", "world"}

	updateSlice(mySlice)
	fmt.Println(mySlice)

}

// To update vaue we most use the * (star) to poit to the value in memory
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}

// Slices works differently the value will become the value we used without pointers
func updateSlice(s []string) {
	s[0] = "World"
}

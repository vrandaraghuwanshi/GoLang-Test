package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	jim := person{
		firstName: "jim",
		lastName:  "party",
		// contact: contactInfo{
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 98890,
		},
	}
	// fmt.Printf("%+v", jim)
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	jim.updateName("jimmy")
	jim.print()
}

// func (p person) updateName(newFirstName string){
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

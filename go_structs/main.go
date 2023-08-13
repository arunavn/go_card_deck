package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstname string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) updateName(newFname string) {
	(*p).firstname = newFname
}

func main() {
	arun := person{firstname: "Arun",
		lastName: "Singh",
		contactInfo: contactInfo{
			"Delhi",
			"+91-5478956416",
		},
	}
	// arun.contact = contactInfo{"Delhi", "+93-5478956416"}
	fmt.Println(arun)
	arun.firstname = "Rohit"
	arun.print()
	// arunPointer := &arun
	// arunPointer.updateName("arun")
	arun.updateName("arun")
	arun.print()

}

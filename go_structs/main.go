package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstname string
	lastName  string
	contact   contactInfo
}

func main() {
	arun := person{firstname: "Arun",
		lastName: "Singh",
		contact: contactInfo{
			"Delhi",
			"+91-5478956416",
		},
	}
	// arun.contact = contactInfo{"Delhi", "+93-5478956416"}
	fmt.Println(arun)
	arun.firstname = "Rohit"
	fmt.Println(arun)
	var arun1 person
	fmt.Printf("%+v \n", arun1)

}

package main

import (
	"log"
	"time"
)

// to get rid of making different variable create a `type`
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

// Receivers example:
type myStruct struct {
	FirstName string;
}

func (m *myStruct) printFirstName() string {
	return m.FirstName;
}

// defining variable like this only available it only in this package
var spcl string

// but if it starts with capital letter it will be available outside the package
// well. Same is true for functions as well
var Spcl string

func main() {
	// take a variable and assign the values in it
	user := User {
		FirstName: "John",
		LastName: "Doe",
	}

	var myVar myStruct;
	myVar.FirstName = "Livo"

	myVar2 := myStruct {
		FirstName: "Hero",
	}

	log.Println(user.FirstName, user.LastName, user.Age, user.PhoneNumber, user.BirthDate)

	log.Println("myVar is ", myVar.printFirstName())
	log.Println("myVar2 is ", myVar2.printFirstName())
}

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

	log.Println(user.FirstName, user.LastName, user.Age, user.PhoneNumber, user.BirthDate)
}

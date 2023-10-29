package main

import "log"

// user defined type
type User struct {
	FirstName string
	LastName string
}
func main() {
//	syntax to define a map
	myMap := make(map[string]string)

	// if don't know what we are going to return as value from map use:
	// interface{}

	myMap["firstName"] = "Go"
	myMap["secondName"] = "Programming language"

	myMap2 := make(map[string]User)

	mySelf := User {
		FirstName: "John",
		LastName: "Doe",
	}

	myMap2["mySelf"] = mySelf

	log.Println(myMap["firstName"], myMap2["mySelf"].FirstName)
}
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


	/// Slices:

	var mySlice []string

	
	mySlice = append(mySlice, "yaten")
	mySlice = append(mySlice, "SHarma")
	mySlice = append(mySlice, "learning!!Go")

	// another approach to declare values in slice
	numbers := []int{1,2,3,4,5,6}
	log.Println((mySlice))
	log.Println(numbers)

	// to print some range of numbers
	log.Println(numbers[0:4])
}
package main

import (
	"fmt"
	"log"
)

var s = "seven"

func main() {
	fmt.Println("Hello world")

	// defining the variables

	var whatToSay string
	var i int

	var s2 = "six"


	whatToSay = "Good bye!!, cruel world"

	fmt.Println(whatToSay)

	i = 100
	fmt.Println("i is set to", i)

	log.Println("s is ", s)
	log.Println("s2 is ", s2)

	// shorthand syntax of variable declration
	whatWasSaid, otherThingSaid := saySomething()

	fmt.Println("saySomething function has returned us", whatWasSaid, otherThingSaid)

	var myStr string
	myStr = "Black"

	log.Println("myStr is set to", myStr)
	changeUsingPointer(&myStr)

	log.Println("After the func call myStr is set to", myStr)

	saySomethingAgain("calling saySomethingAgain")

}

// syntax of function i.e this function will going to return a `string`
func saySomething() (string, string) {
	// speciality of goLang, we can return multiple things from a single func.
	return "something", "else"
}

// case of pointer and using it: use `&` for reference and `*` for value
func changeUsingPointer(s *string) {
	newValue := "White"
	*s = newValue
}

// if i pass same param as `global variable` so when i will call this function it
// will refer to that variable.
func saySomethingAgain(s3 string) (string, string) {
	log.Println("s from saySomethingAgain is ", s)
	return s, "returned";
}
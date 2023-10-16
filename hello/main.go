package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello world")

	// defining the variables

	var whatToSay string
	var i int


	whatToSay = "Good bye!!, cruel world"

	fmt.Println(whatToSay)

	i = 100
	fmt.Println("i is set to", i)

	// shorthand syntax of variable declration
	whatWasSaid, otherThingSaid := saySomething()

	fmt.Println("saySomething function has returned us", whatWasSaid, otherThingSaid)

	var myStr string
	myStr = "Black"

	log.Println("myStr is set to", myStr)
	changeUsingPointer(&myStr)

	log.Println("After the func call myStr is set to", myStr)

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
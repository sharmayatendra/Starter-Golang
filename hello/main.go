package main

import (
	"log"
)

var s = "seven"

func main() {
	// concept & syntax of if/else, switch statement:

	myNum := 100
	isTrue := true

	// test with two variable conditions

	if myNum > 100 && isTrue {
		log.Println("myNum is greater than 100")
	} else if myNum == 100 && !isTrue {
		log.Println("myNum is equals to 100")
	} else {
		log.Println("myNum is lesser than 100")
	}





}
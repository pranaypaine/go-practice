package main

import "fmt"

func main() {
	// Init local variable
	var v1 int = 700

	//Checking the condition
	if v1 == 100 {
		// if condition is true then
		// display the following
		fmt.Printf("value of v1 is 100\n")
	} else if v1 == 200 {
		fmt.Printf("value of v1 is 200\n")
	} else if v1 == 300 {
		fmt.Printf("value of v1 is 300\n")
	} else {
		// if none of the conditions is true
		fmt.Printf("None of the value are matching")
	}

}

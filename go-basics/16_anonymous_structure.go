package main

import "fmt"

// Creating a structure
// with anonymous fields
type Student struct {
	int
	string
	float64
}

func main() {
	// Creating and initializing
	// the anonymous structure
	Element := struct {
		Name string
	}{
		Name: "PikaChu",
	}

	// Display the anonymous structure
	fmt.Println(Element)

	// Assigning values to the anonymous
	// fields of the student structure
	value := Student{123, "Bud", 89.89}

	// Display the values of the fields
	fmt.Println("Enrollment number : ", value.int)
	fmt.Println("Student name : ", value.string)
	fmt.Println("Package price : ", value.float64)

}

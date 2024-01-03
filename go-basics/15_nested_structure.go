package main

import "fmt"

// Creating structure
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// Creating nested structure
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address
}

func main() {
	fmt.Println("Nested Structures")

	// Initializing the fields
	// of the structure
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address: Address{
			Street:     "123 Main St",
			City:       "Any Town",
			State:      "CA",
			PostalCode: "12345",
		},
	}

	// Display the values
	fmt.Println("Name:", p.FirstName, p.LastName)
	fmt.Println("Age:", p.Age)
	fmt.Println("Address:")
	fmt.Println("Street:", p.Address.Street)
	fmt.Println("City:", p.Address.City)
	fmt.Println("State:", p.Address.State)
	fmt.Println("Postal Code:", p.Address.PostalCode)
}

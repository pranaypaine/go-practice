package main

import "fmt"

type Address struct {
	Name    string
	City    string
	Pincode int
}

func main() {
	fmt.Println("Structures")

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a Address
	fmt.Println(a)

	// Declaring and initializing a
	// struct using a struct literal
	a1 := Address{"akshay", "dehradun", 3623752}
	fmt.Println("Address 1:", a1)

	// Naming fields while
	// initializing a struct
	a2 := Address{Name: "Ankita", City: "Ballia", Pincode: 277001}
	fmt.Println("Address 2:", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := Address{City: "Delhi"}
	fmt.Println("Address 3:", a3)

	// Accessing struct fields
	// using the dot operator
	fmt.Println("Name:", a2.Name)
	fmt.Println("city:", a2.City)

	// Assigning a new value
	// to a struct field
	a2.City = "Gurugram"
	fmt.Println("city:", a2.City)

	// passing the address of struct variable
	// a4 is a pointer to the Address struct
	a4 := &Address{Name: "Shivam", City: "Noida", Pincode: 2323223}

	// (*a4).Name is the syntax to access
	// the Name field of the a4 struct
	fmt.Println("Name:", (*a4).Name)

	// a4.Name is used to access
	// the field Name
	fmt.Println("Name:", a4.Name)
}

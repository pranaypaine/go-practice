package main

import "fmt"

// Author Structure
type author struct {
	first_name      string
	last_name       string
	books_published int
}

func (a author) show() {
	fmt.Println("Author First Name: ", a.first_name)
	fmt.Println("Author Last Name: ", a.last_name)
	fmt.Println("Books Published :", a.books_published)
}

func main() {
	fmt.Println("Methods")

	// initializing the values
	// of authorr sructure
	res := author{
		first_name:      "Pranay",
		last_name:       "Paine",
		books_published: 10,
	}
	// Calling the method
	res.show()
}

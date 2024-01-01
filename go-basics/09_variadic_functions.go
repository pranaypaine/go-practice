package main

import (
	"fmt"
	"strings"
)

func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	// No Argumants
	fmt.Println(joinstr())

	// Multiple Arguments
	fmt.Println(joinstr("GEEK", "GFG"))
	fmt.Println(joinstr("Geeks", "For", "Geeks"))
	fmt.Println(joinstr("G", "E", "E", "k", "S"))

	// pass a slice in variadic function

	elements := []string{"Geeks", "For", "Geeks"}
	fmt.Println(joinstr(elements...))
}

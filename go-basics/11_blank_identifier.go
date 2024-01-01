// Blank identifier helps in function that return multiple values
// but we need some specific value,
// to ignore the extra value we use _ to skip the value we don't need

package main

import "fmt"

func main() {
	// calling the function
	// function returns two values which are
	// assigned to mul and blank identifier
	mul, _ := mul_div(105, 7)

	// only using the mul variable
	fmt.Println("105 x 7 = ", mul)
}

// function returning two
// values of integer type
func mul_div(n1 int, n2 int) (int, int) {

	// returning the values
	return n1 * n2, n1 / n2
}

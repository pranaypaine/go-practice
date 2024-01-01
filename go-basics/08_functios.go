package main

import "fmt"

// Function declaration
// func function_name(Parameter List)(Return Type)  {
// 	Function Body
// }

func calculate_area(length, breadth int) int {
	area := length * breadth
	return area
}

func main() {
	fmt.Println("Funtions")
	fmt.Printf("Area of rectangle is : %d", calculate_area(10, 12))
}

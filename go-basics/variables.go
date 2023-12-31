package main

import "fmt"

func main() {
	// Variable declared and
	// initialized without the
	// explicit type
	var myvariable1 = 20
	var myvariable2 = "GeeksforGeeks"
	var myvariable3 = 34.80

	// Display the value and the
	// type of the variables
	fmt.Printf("The value of myvariable1 is : %d\n",
		myvariable1)

	fmt.Printf("The type of myvariable1 is : %T\n",
		myvariable1)

	fmt.Printf("The value of myvariable2 is : %s\n",
		myvariable2)

	fmt.Printf("The type of myvariable2 is : %T\n",
		myvariable2)

	fmt.Printf("The value of myvariable3 is : %f\n",
		myvariable3)

	fmt.Printf("The type of myvariable3 is : %T\n",
		myvariable3)

	// Variable declared and
	// initialized without expression
	var myvariable4 int
	var myvariable5 string
	var myvariable6 float64

	// Display the zero-value of the variables
	fmt.Printf("The value of myvariable4 is : %d\n",
		myvariable4)

	fmt.Printf("The value of myvariable5 is : %s\n",
		myvariable5)

	fmt.Printf("The value of myvariable6 is : %f\n",
		myvariable6)

	// Multiple variables of the same type
	// are declared and initialized
	// in the single line
	var myvariable7, myvariable8, myvariable9 int = 2, 454, 67

	// Display the values of the variables
	fmt.Printf("The value of myvariable7 is : %d\n",
		myvariable7)

	fmt.Printf("The value of myvariable8 is : %d\n",
		myvariable8)

	fmt.Printf("The value of myvariable9 is : %d\n",
		myvariable9)

	// Multiple variables of different types
	// are declared and initialized in the single line
	var myvariable10, myvariable11, myvariable12 = 2, "GFG", 67.56

	// Display the value and
	// type of the variables
	fmt.Printf("The value of myvariable10 is : %d\n",
		myvariable10)

	fmt.Printf("The type of myvariable10 is : %T\n",
		myvariable10)

	fmt.Printf("The value of myvariable11 is : %s\n",
		myvariable11)

	fmt.Printf("The type of myvariable11 is : %T\n",
		myvariable11)

	fmt.Printf("The value of myvariable12 is : %f\n",
		myvariable12)

	fmt.Printf("The type of myvariable12 is : %T\n",
		myvariable12)

	// Using short variable declaration
	myvar1 := 39
	myvar2 := "GeeksforGeeks"
	myvar3 := 34.67

	// Display the value and type of the variables
	fmt.Printf("The value of myvar1 is : %d\n", myvar1)
	fmt.Printf("The type of myvar1 is : %T\n", myvar1)

	fmt.Printf("The value of myvar2 is : %s\n", myvar2)
	fmt.Printf("The type of myvar2 is : %T\n", myvar2)

	fmt.Printf("The value of myvar3 is : %f\n", myvar3)
	fmt.Printf("The type of myvar3 is : %T\n", myvar3)

	// Using short variable declaration
	// Multiple variables of same types
	// are declared and initialized in
	// the single line
	myvar4, myvar5, myvar6 := 800, 34, 56

	// Display the value and
	// type of the variables
	fmt.Printf("The value of myvar4 is : %d\n", myvar4)
	fmt.Printf("The type of myvar4 is : %T\n", myvar4)

	fmt.Printf("The value of myvar5 is : %d\n", myvar5)
	fmt.Printf("The type of myvar5 is : %T\n", myvar5)

	fmt.Printf("The value of myvar6 is : %d\n", myvar6)
	fmt.Printf("The type of myvar6 is : %T\n", myvar6)

	// Using short variable declaration
	// Multiple variables of different types
	// are declared and initialized in the single line
	myvar7, myvar8, myvar9 := 800, "Geeks", 47.56

	// Display the value and type of the variables
	fmt.Printf("The value of myvar7 is : %d\n", myvar7)
	fmt.Printf("The type of myvar7 is : %T\n", myvar7)

	fmt.Printf("The value of myvar8 is : %s\n", myvar8)
	fmt.Printf("The type of myvar8 is : %T\n", myvar8)

	fmt.Printf("The value of myvar9 is : %f\n", myvar9)
	fmt.Printf("The type of myvar9 is : %T\n", myvar9)

}

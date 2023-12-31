package main

import "fmt"

func main() {

	//basic Arithmetic operations
	p := 34
	q := 20

	// Addition
	result1 := p + q
	fmt.Printf("Result of p + q = %d", result1)

	// Subtraction
	result2 := p - q
	fmt.Printf("\nResult of p - q = %d", result2)

	// Multiplication
	result3 := p * q
	fmt.Printf("\nResult of p * q = %d", result3)

	// Division
	result4 := p / q
	fmt.Printf("\nResult of p / q = %d", result4)

	// Modulus
	result5 := p % q
	fmt.Printf("\nResult of p %% q = %d", result5)

	// Rational Operators
	// ‘=='(Equal To)
	result6 := p == q
	fmt.Println(result6)

	// ‘!='(Not Equal To)
	result7 := p != q
	fmt.Println(result7)

	// ‘<‘(Less Than)
	result8 := p < q
	fmt.Println(result8)

	// ‘>'(Greater Than)
	result9 := p > q
	fmt.Println(result9)

	// ‘>='(Greater Than Equal To)
	result10 := p >= q
	fmt.Println(result10)

	// ‘<='(Less Than Equal To)
	result11 := p <= q
	fmt.Println(result11)

	// Logical Operators
	if p != q && p <= q {
		fmt.Println("True")
	}

	if p != q || p <= q {
		fmt.Println("True")
	}

	if !(p == q) {
		fmt.Println("True")
	}
}

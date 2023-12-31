package main

import (
	"fmt"
	"math"
)

// Single Constant without type
const PI = 3.14

// Single Connstant with type
const s string = "GeeksForGeeks"

// Multiple constants declared at once
const (
	GFG     = "GeeksforGeeks"
	Correct = true
	Pi      = 3.14
)

func main() {
	const GFG = "GeeksforGeeks"
	fmt.Println("hello", GFG)
	fmt.Println("Happy", PI, "Day")

	const Correct = true
	fmt.Println("Go rules?", Correct)

	fmt.Println(s)

	const n = 5

	const d = 3e10 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

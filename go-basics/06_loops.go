package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 4; i++ {
		fmt.Printf("This is a for loop\n")
	}

	// while loop
	// for loop executes till
	// i < 3 condition is true
	i := 0

	for i < 3 {
		i += 2
		print(i)
	}
	fmt.Println(i)

	// Loop using range variable
	rangeVariable := []string{"GFG", "Geeks", "GeeksforGeeks"}

	for i, j := range rangeVariable {

		fmt.Println(i, j)
	}
	// String as a range in the for loop
	for i, j := range "XabCd" {
		fmt.Printf("The index number of %U is %d\n", j, i)
	}

	// loops using maps for extracting key value pairs
	newmap := map[int]string{
		22: "Geeks",
		33: "GFG",
		44: "GeeksForGeeks",
	}

	for key, value := range newmap {
		fmt.Println(key, value)
	}

	// using channel
	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i := range chnl {
		fmt.Println(i)
	}
}

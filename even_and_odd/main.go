package main

import "fmt"

func main() {
	// Option #1: Create slice of integers manually
	//integers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Option #2: Create slice of integers using a for loop
	var integers []int
	for i := 0; i <= 10; i++ {
		integers = append(integers, i)
	}

	// Loop through the slice of integers
	for _, value := range integers {
		// If value mod 2 == 0 then its an even number else its an odd number
		if value%2 == 0 {
			fmt.Println(value, " is even")
		} else {
			fmt.Println(value, " is odd")
		}
	}
}

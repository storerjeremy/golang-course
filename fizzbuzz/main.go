package main

import "fmt"

// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print “Fizz” instead of the number
// and for the multiples of five print “Buzz”. For numbers which are
// multiples of both three and five print “FizzBuzz”."

func main() {
	for i := 1; i <= 100; i++ {
		fizzBuzz(i)
	}
}

func fizzBuzz(n int) {
	if n%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(n)
	}
}

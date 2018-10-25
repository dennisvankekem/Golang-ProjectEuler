package main

import (
	"fmt"
)

func main() {

	//The prime factors of 13195 are 5, 7, 13 and 29.
	//What is the largest prime factor of the number 600851475143 ?

	fmt.Println(largestPrimeFactor(600851475143))
}

// largestPrimeFactor checks for largest prime number in given number
func largestPrimeFactor(n int) int {

	x := 2
	for n > 1 {
		if n%x == 0 {
			n = n / x
		} else {
			x++
		}
	}

	return x
}

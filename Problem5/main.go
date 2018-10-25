package main

import (
	"fmt"
)

func main() {
	//2520 is the smallest number that can be divided by each of the numbers from 1 to 10
	//without any remainder.
	//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

	fmt.Println(smallestNumberEvenlyDivisible(20))
}

func smallestNumberEvenlyDivisible(b int) int {

	n := 1

	for i := 1; i <= b; {
		if n%i == 0 {
			i++
		} else {
			n++
			i = 1
		}
	}

	return n
}

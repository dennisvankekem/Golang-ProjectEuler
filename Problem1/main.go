package main

import (
	"fmt"
)

func main() {

	//If we list all the natural numbers below 10 that are multiples of 3 or 5,
	// we get 3, 5, 6 and 9. The sum of these multiples is 23.
	//Find the sum of all the multiples of 3 or 5 below 1000.

	fmt.Println(sumOfMultiples(1000, 3, 5))
}

func sumOfMultiples(baseNumber int, lowMultiple int, highMultiple int) int {

	totalCount := 0

	for i := 0; i < baseNumber; i++ {
		fizz := i%lowMultiple == 0
		buzz := i%highMultiple == 0

		if fizz && buzz {
			totalCount += i
		} else if fizz {
			totalCount += i
		} else if buzz {
			totalCount += i
		}

	}

	return totalCount
}

package main

import (
	"fmt"
)

func main() {
	//The sum of the squares of the first ten natural numbers is,
	//1sq + 2sq + ... + 10sq = 385
	//The square of the sum of the first ten natural numbers is,
	//(1 + 2 + ... + 10)sq = 55sq = 3025
	//Hence the difference between the sum of the squares of the first ten natural numbers
	//and the square of the sum is 3025 − 385 = 2640.
	//Find the difference between the sum of the squares of the first one hundred natural numbers
	//and the square of the sum.
	x := 100
	fmt.Println(squareOfSum(x) - sumOfSquares(x))

}

func sumOfSquares(s int) int {

	n := 0

	for i := 1; i <= s; i++ {
		n += i * i
	}

	return n
}

func squareOfSum(s int) int {

	n := 0

	for i := 1; i <= s; i++ {
		n += i
	}

	n = n * n

	return n
}

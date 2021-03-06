package main

import (
	"fmt"
)

func main() {
	//Each new term in the Fibonacci sequence is generated by adding the previous two terms.
	//By starting with 1 and 2, the first 10 terms will be:

	//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

	//By considering the terms in the Fibonacci sequence whose values do not exceed four million,
	//find the sum of the even-valued terms.

	fmt.Println(Fibonacci(100))

}

//Fibonacci returns sum of equals in Fibonacci sequence with limit = l
func Fibonacci(l int) int {
	s := 0
	x := 1
	y := 1
	z := 0
	fmt.Println(x)
	for x < l {
		fmt.Println(x)
		if x%2 == 0 {
			s += x
		}
		z = x
		x += y
		y = z
	}
	return s
}

package main

import (
	"fmt"
)

func main() {
	//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
	//we can see that the 6th prime is 13.
	//What is the 10 001st prime number?

	fmt.Println(OptimusPrime(10001))
}

// OptimusPrime gets all the prime numbers and returns primenumber on given position
func OptimusPrime(n int) int {

	c := 1
	x := 3
	s := []int{2}
	b := false

	for c <= n {

		for _, value := range s {
			if x%value == 0 {
				x++
				b = false
				break
			}
			b = true
		}

		if b {
			s = append(s, x)
			c++
			x++
		}
		b = false
	}

	return s[n-1]

}

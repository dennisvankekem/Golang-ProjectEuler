package main

import (
	"fmt"
)

func main() {
	//A palindromic number reads the same both ways.
	//The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
	//Find the largest palindrome made from the product of two 3-digit numbers.

	fmt.Println(largestPalindromeNumber(999, 999))

}

func largestPalindromeNumber(s int, l int) int {
	b := 0
	lb := 0
	for i := 1; i <= s; i++ {
		for j := 1; j <= l; j++ {
			b = i * j
			bRef := mathReverseNumber(b)
			if b == bRef {
				if b > lb {
					lb = b
				}
			}
		}
	}
	return lb
}

func mathReverseNumber(n int) int {

	r := 0
	for n > 0 {
		l := n % 10
		r = (r * 10) + l
		n = n / 10
	}

	return r
}

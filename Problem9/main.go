package main

import (
	"fmt"
	"math"
)

func main() {
	//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
	//a sr + b sr = c sr
	//For example, 3 sr + 4 sr = 9 + 16 = 25 = 5 sr.
	//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	//Find the product abc.

	isPythagoreanTriplet(1000)
}

func isPythagoreanTriplet(l int) {
	a := 1
	b := 1
	m := l / 2

	for i := 1; i < m; i++ {
		for j := 1; j < m; j++ {
			if i > j {
				a = i
				b = j
				c := math.Sqrt(float64((a * a) + (b * b)))
				if c == float64(int64(c)) {
					d := a + b + int(c)
					if d == l {
						fmt.Println("a = ", a, ", b = ", b, ", c = ", c)
					}
				}
			}
		}
	}

}

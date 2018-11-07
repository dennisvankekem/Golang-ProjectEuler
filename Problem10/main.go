package main

import "fmt"

func main() {
	//yes, lazy I know. Basicly problem7 extreme
	OptimusPrime(2000000)

}

// OptimusPrime gets all the prime numbers and returns primenumber addition until limit is reached
func OptimusPrime(n int) int {

	c := 2
	x := 3
	s := []int{2}
	b := false

	for x <= n {

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
			c += x
			x++
		}
		b = false
	}

	fmt.Println(c)
	return c

}

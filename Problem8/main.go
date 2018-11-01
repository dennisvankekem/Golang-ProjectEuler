package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	//The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.
	//Find the thirteen adjacent digits in the 1000-digit number that have the greatest product.
	//What is the value of this product?

	//amount of adjacent digits to be matched
	a := 13
	//read file with numbers (expects formatted document, no error handeling here)
	s := readFile()
	//converts file to slice of ints
	f, err := SliceOfByteToSliceOfInt(strings.NewReader(s))
	checkForError(err)
	//peforms the formula on set of numbers with given context
	fmt.Println(greatestNumberInAdjacentDigits(a, f))
}

//checkForError handles general errors
func checkForError(e error) {
	if e != nil {
		panic(e)
	}
}

//readFile reads a file and returns a stringified byteslice
func readFile() string {
	f, err := ioutil.ReadFile("number")
	checkForError(err)
	s := string(f)

	return s
}

//SliceOfByteToSliceOfInt takes a byteslice and converts it to slice of int
func SliceOfByteToSliceOfInt(r io.Reader) ([]int, error) {

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	var s []int
	for scanner.Scan() {
		if scanner.Text() != "\n" {
			d, err := strconv.Atoi(scanner.Text())
			checkForError(err)
			s = append(s, d)
		}
	}
	return s, scanner.Err()
}

//greatestNumberInAdjacentDigits does just that with given count of numbers to check
func greatestNumberInAdjacentDigits(c int, ls []int) int {

	y := 1
	x := 1
	j := 0

	for j < len(ls)-c {
		for i := 0; i < c; i++ {
			x = x * ls[j+i]
		}
		if x > y {
			y = x
		}
		x = 1
		j++
	}

	return y
}

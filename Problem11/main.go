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
	l := 4
	s := readFile()
	f, err := SliceOfByteToSliceOfIntPaired(strings.NewReader(s))
	checkForError(err)
	fmt.Println(calculateBiggestMultiple(f, l))

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

//SliceOfByteToSliceOfIntPaired takes a byteslice and converts it to slice of int in pairs Split by whitespace
func SliceOfByteToSliceOfIntPaired(r io.Reader) ([][]int, error) {

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var s []int
	var ms [][]int
	i := 0
	for scanner.Scan() {
		if i != 0 && i%20 == 0 {
			ms = append(ms, s)
			s = nil
		}
		d, err := strconv.Atoi(scanner.Text())
		checkForError(err)
		s = append(s, d)
		i++
	}

	ms = append(ms, s)

	return ms, scanner.Err()
}

func calculateBiggestMultiple(nf [][]int, l int) int {

	bn := 0

	//horizontal
	for i := 0; i < 20; i++ {
		for j := 0; j < 20-l; j++ {
			nr := nf[i][j] * nf[i][j+1] * nf[i][j+2] * nf[i][j+3]
			if nr > bn {
				bn = nr
			}
		}
	}

	//vertical
	for i := 0; i < 20-l; i++ {
		for j := 0; j < 20; j++ {
			nr := nf[i][j] * nf[i+1][j] * nf[i+2][j] * nf[i+3][j]
			if nr > bn {
				bn = nr
			}
		}
	}

	//diagonal positive
	for i := 0; i < 20-l+1; i++ {
		for j := 0; j < 20-l+1; j++ {
			nr := nf[i][j] * nf[i+1][j+1] * nf[i+2][j+2] * nf[i+3][j+3]
			if nr > bn {
				bn = nr
			}
		}
	}

	//diagonal negative
	for i := 0; i < 20-l+1; i++ {
		for j := 0 + l - 1; j < 20; j++ {
			nr := nf[i][j] * nf[i+1][j-1] * nf[i+2][j-2] * nf[i+3][j-3]
			if nr > bn {
				bn = nr
			}
		}
	}

	return bn
}

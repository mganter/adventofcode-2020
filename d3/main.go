package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	i, _ := ioutil.ReadFile("values")
	input := to2d(string(i))

	fmt.Printf("Part ONE: \n")
	slopeStyle := SlopeStype{1, 3}
	fmt.Printf("%v\n", Slope(input, slopeStyle))

	fmt.Printf("Part TWO: \n")
	slopes := []SlopeStype{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	prod := 1
	for _, slope := range slopes {
		prod *= Slope(input, slope)
	}
	fmt.Printf("%v\n", prod)
}

func Slope(field [][]bool, slopeStyle SlopeStype) int {
	currHor := 0
	currVer := 0
	counter := 0

	for ; currVer < len(field); currHor %= len(field[0]) {
		if field[currVer][currHor] {
			counter++
		}
		currVer += slopeStyle.down
		currHor += slopeStyle.right
	}
	return counter
}

type SlopeStype struct {
	down  int
	right int
}

func to2d(text string) [][]bool {
	arr := make([][]bool, 0)
	for _, line := range strings.Split(text, "\n") {
		row := make([]bool, 0)
		for index := range line {
			row = append(row, line[index] == '#')
		}
		arr = append(arr, row)
	}
	return arr
}

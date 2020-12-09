package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	vals := GetValues()

	counter := 0
	for _, val := range vals {
		if checkPassword(val) {
			counter++
		}
	}

	fmt.Printf("%v\n", counter)

	counter2 := 0
	for _, val := range vals {
		if checkPassword2(val) {
			counter2++
		}
	}

	fmt.Printf("%v\n", counter2)
}

func checkPassword(passwordAndPolicy string) bool {
	minMax, requiredChar, password := extractPolicy(passwordAndPolicy)
	occurrences := strings.Count(password, requiredChar)
	return occurrences >= minMax.min && occurrences <= minMax.max
}

func checkPassword2(passwordAndPolicy string) bool {
	minMax, requiredChar, password := extractPolicy(passwordAndPolicy)
	return (password[minMax.min-1] == requiredChar[0]) != (password[minMax.max-1] == requiredChar[0])
}

func extractPolicy(passwordAndPolicy string) (Range, string, string) {
	arr := strings.Split(passwordAndPolicy, " ")
	password := arr[2]
	character := strings.TrimRight(arr[1], ":")

	r := strings.Split(arr[0], "-")
	min, _ := strconv.Atoi(r[0])
	max, _ := strconv.Atoi(r[1])

	return Range{min, max}, character, password
}

type Range struct {
	min int
	max int
}
